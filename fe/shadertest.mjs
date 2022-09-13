import { Vector2D } from './Libs/Math/vector.mjs'
import * as mat4 from './src/mat4.mjs'
import * as vec2 from './src/vec2.mjs'

const canvas = document.querySelector('canvas')

// const m = canvas.getContext("2d");
// m.fillStyle = "blue";
// m.fillRect(0, 0, 800, 800);
main();

//
// Start here
//
function main() {
  const gl = canvas.getContext('webgl');
  // If we don't have a GL context, give up now

  if (!gl) {
    alert('Unable to initialize WebGL. Your browser or machine may not support it.');
    return;
  }

  // Vertex shader program

  const vsSource = `

    uniform mat2 u_transformRotation;
    uniform vec2 u_transformScale;
    uniform vec2 u_transformPosSubCamPos;
    uniform vec2 u_cameraSizeInverted;
  
    attribute vec2 a_vertexPosition;
    attribute vec2 a_vertexUV;
    attribute vec4 a_color;

    varying  vec4 v2f_color;
    varying  vec2 v2f_UV;

    void main() {

      vec2 temp = (u_transformRotation * (a_vertexPosition * u_transformScale) + u_transformPosSubCamPos) * u_cameraSizeInverted;
      gl_Position = vec4(temp.x, temp.y, 0, 1);
      v2f_color = a_color;
      v2f_UV = a_vertexUV;
    }
  `;

  // Fragment shader program

  const fsSource = `
    precision mediump float;

    #define SQ(X) ((X)*(X))
    #define CompareDistanceLess(P1,P2,D) (SQ((P1).x-(P2).x)+SQ((P1).y-(P2).y)<SQ(D))

    // Require resolution (canvas size) as an input
    //uniform vec3 uResolution;
    uniform sampler2D u_image;

    varying  vec4 v2f_color;
    varying  vec4 v2f_UV;

    void main() {

      // Calculate relative coordinates (uv)
    //   vec2 uv = gl_FragCoord.xy / uResolution.xy;
    //   vec2 point = vec2(0.69,0.5);
    //   point -= vec2(0.5,0.5);
    //   point = vec2(point.x,-point.y);
    //   point = vec2(point.y,point.x);
    //   point += vec2(0.5,0.5);
      vec2 vec2color = v2f_UV.xy;
    //   if (CompareDistanceLess(uv,point,0.1)){
    //     vec2color = vec2(1,1);
    //   }
    //   else{
    //     vec2color = vec2(1,0);
    //   }
      gl_FragColor = vec4(vec2color.x,vec2color.y,0,1);
    //   gl_FragColor = v2f_color;
    //   gl_FragColor = vec4(uv.x, uv.y, 0,1);

      gl_FragColor = texture2D(u_image, v2f_UV) * f_;
    }
  `;

  // Initialize a shader program; this is where all the lighting
  // for the vertices and so forth is established.
  const shaderProgram = initShaderProgram(gl, vsSource, fsSource);

  // Collect all the info needed to use the shader program.
  // Look up which attribute our shader program is using
  // for a_vertexPosition and look up uniform locations.
  const programInfo = {
    program: shaderProgram,
    attribLocations: {
      vertexPosition: gl.getAttribLocation(shaderProgram, 'a_vertexPosition'),
      vertexUV: gl.getAttribLocation(shaderProgram, 'a_vertexUV'),
      vertexColor: gl.getAttribLocation(shaderProgram, 'a_color'),
    },
    uniformLocations: {
      transformRotaion: gl.getUniformLocation(shaderProgram, 'u_transformRotation'),
      transformScale: gl.getUniformLocation(shaderProgram, 'u_transformScale'),
      transformPosSubCamPos: gl.getUniformLocation(shaderProgram, 'u_transformPosSubCamPos'),
      cameraSizeInverted: gl.getUniformLocation(shaderProgram, 'u_cameraSizeInverted'),
    },
  };

  // Here's where we call the routine that builds all the
  // objects we'll be drawing.

  const positions = [
    1, 1,
    -1, 1,
    1, -1,
    -1, -1,
  ];

  const uv = [
    1, 1,
    0, 1,
    1, 0,
    0, 0,
  ];

  const colors = [
    0, 1, 0, 1,
    1, 0, 0, 1,
    0, 0, 1, 1,
    1, 1, 0, 1,
  ];

  const data = {
    positions: positions,
    uv: uv,
    colors: colors,
  };

  const buffers = initBuffers(gl, data);


  const transformPosition = Vector2D(10, 10)
  const transformScale = Vector2D(1, 1)
  const transformRotation = Vector2D(1, 0)

  const objectData = {
    position: transformPosition,
    scale: transformScale,
    rotation: transformRotation,
  }

  const cameraPos = Vector2D(0, 0)
  const cameraSizeInverted = Vector2D(1 / 40, 1 / 40)

  const cameraData = {
    position: cameraPos,
    sizeInverted: cameraSizeInverted,
  }

  // Draw the scene
  drawEntity(gl, programInfo, buffers, objectData, cameraData);
}

//
// initBuffers
//
// Initialize the buffers we'll need. For this demo, we just
// have one object -- a simple two-dimensional square.
//
function initBuffers(gl, data) {

  // Create a buffer for the square's positions.

  const positionBuffer = gl.createBuffer();
  const uvBuffer = gl.createBuffer();
  const colorBuffer = gl.createBuffer();

  // Now create an array of positions for the square.

  // Select the positionBuffer as the one to apply buffer
  // operations to from here out.

  gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);

  // Now pass the list of positions into WebGL to build the
  // shape. We do this by creating a Float32Array from the
  // JavaScript array, then use it to fill the current buffer.



  gl.bufferData(gl.ARRAY_BUFFER,
    new Float32Array(data.positions),
    gl.STATIC_DRAW);


  gl.bindBuffer(gl.ARRAY_BUFFER, uvBuffer);
  gl.bufferData(gl.ARRAY_BUFFER,
    new Float32Array(data.uv),
    gl.STATIC_DRAW);


  gl.bindBuffer(gl.ARRAY_BUFFER, colorBuffer);
  gl.bufferData(gl.ARRAY_BUFFER,
    new Float32Array(data.colors),
    gl.STATIC_DRAW);

  return {
    position: positionBuffer,
    uv: uvBuffer,
    color: colorBuffer,
  };
}

//
// Draw the scene.
//
function drawEntity(gl, programInfo, buffers, objectData, cameraData) {
  gl.clearColor(0.0, 0.0, 0.0, 1.0);  // Clear to black, fully opaque
  gl.clearDepth(1.0);                 // Clear everything
  gl.enable(gl.DEPTH_TEST);           // Enable depth testing
  gl.depthFunc(gl.LEQUAL);            // Near things obscure far things

  // Clear the canvas before we start drawing on it.

  gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

  // Tell WebGL how to pull out the positions from the position
  // buffer into the vertexPosition attribute.
  {
    const numComponents = 2;
    const type = gl.FLOAT;
    const normalize = false;
    const stride = 0;
    const offset = 0;
    gl.bindBuffer(gl.ARRAY_BUFFER, buffers.position);
    gl.vertexAttribPointer(
      programInfo.attribLocations.vertexPosition,
      numComponents,
      type,
      normalize,
      stride,
      offset);
    gl.enableVertexAttribArray(
      programInfo.attribLocations.vertexPosition);
  }

  {
    const numComponents = 2;
    const type = gl.FLOAT;
    const normalize = false;
    const stride = 0;
    const offset = 0;
    gl.bindBuffer(gl.ARRAY_BUFFER, buffers.uv);
    gl.vertexAttribPointer(
      programInfo.attribLocations.vertexUV,
      numComponents,
      type,
      normalize,
      stride,
      offset);
    gl.enableVertexAttribArray(
      programInfo.attribLocations.vertexUV);
  }

  {
    const numComponents = 4;
    const type = gl.FLOAT;
    const normalize = false;
    const stride = 0;
    const offset = 0;
    gl.bindBuffer(gl.ARRAY_BUFFER, buffers.color);
    gl.vertexAttribPointer(
      programInfo.attribLocations.vertexColor,
      numComponents,
      type,
      normalize,
      stride,
      offset);
    gl.enableVertexAttribArray(
      programInfo.attribLocations.vertexColor);
  }

  // Tell WebGL to use our program when drawing

  gl.useProgram(programInfo.program);

  // Set the shader uniforms
  let m
  let v
  
  m = mat4.create()
  m[0] = objectData.rotation.x
  m[1] = -objectData.rotation.y
  m[2] = objectData.rotation.y
  m[3] = objectData.rotation.x
  gl.uniformMatrix4fv(
    programInfo.uniformLocations.transformRotaion,
    false,
    m);

  v = vec2.create()
  v[0] = objectData.scale.x
  v[1] = objectData.scale.y
  gl.uniformMatrix4fv(
    programInfo.uniformLocations.transformScale,
    false,
    v);

  v = vec2.create()
  v[0] = objectData.position.x - cameraData.position.x
  v[1] = objectData.position.y - cameraData.position.y
  gl.uniformMatrix4fv(
    programInfo.uniformLocations.transformPosSubCamPos,
    false,
    v);

  v = vec2.create()
  v[0] = cameraData.sizeInverted.x
  v[1] = cameraData.sizeInverted.y
  gl.uniformMatrix4fv(
    programInfo.uniformLocations.cameraSizeInverted,
    false,
    v);


  gl.uniform3f(programInfo.uniformLocations.resolution, canvas.width, canvas.height, 1.0);

  {
    const offset = 0;
    const vertexCount = 4;
    gl.drawArrays(gl.TRIANGLE_STRIP, offset, vertexCount);
  }
}

//
// Initialize a shader program, so WebGL knows how to draw our data
//
function initShaderProgram(gl, vsSource, fsSource) {
  const vertexShader = loadShader(gl, gl.VERTEX_SHADER, vsSource);
  const fragmentShader = loadShader(gl, gl.FRAGMENT_SHADER, fsSource);

  // Create the shader program

  const shaderProgram = gl.createProgram();
  gl.attachShader(shaderProgram, vertexShader);
  gl.attachShader(shaderProgram, fragmentShader);
  gl.linkProgram(shaderProgram);

  // If creating the shader program failed, alert

  if (!gl.getProgramParameter(shaderProgram, gl.LINK_STATUS)) {
    alert('Unable to initialize the shader program: ' + gl.getProgramInfoLog(shaderProgram));
    return null;
  }

  return shaderProgram;
}

//
// creates a shader of the given type, uploads the source and
// compiles it.
//
function loadShader(gl, type, source) {
  const shader = gl.createShader(type);

  // Send the source to the shader object

  gl.shaderSource(shader, source);

  // Compile the shader program

  gl.compileShader(shader);

  // See if it compiled successfully

  if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
    alert('An error occurred compiling the shaders: ' + gl.getShaderInfoLog(shader));
    gl.deleteShader(shader);
    return null;
  }

  return shader;
}