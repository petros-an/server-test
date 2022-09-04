m = document.getElementById("life").getContext("2d");
//alert(typeof m)
const width = 800
const height = 800

const draw = (x, y, c, w, h) => {
  m.fillStyle = c;
  m.fillRect(x, y, w, h);
};


atoms = [];
const atom = (x, y, c) => {
  return { x: x, y: y, vx: 0, vy: 0, color: c };
};


const random = (min, max) => {
    return Math.random() * (max - min) + min;
};


const create = (number, color) => {
  group = [];
  for (let i = 0; i < number; i++) {
    group.push(atom(random(50, width-50), random(50, height-50), color));
    atoms.push(group[i]);
  }
  return group;
};


const rule = (atoms1, atoms2, g) => {
  for (let i = 0; i < atoms1.length; i++) {
    fx = 0;
    fy = 0;
    for (let j = 0; j < atoms2.length; j++) {
      a = atoms1[i];
      b = atoms2[j];
      dx = a.x - b.x;
      dy = a.y - b.y;
      d = Math.sqrt(dx * dx + dy * dy);
      if (d > 0 && d < 80) {
        F = (g * 1) / d;
        fx += F * dx;
        fy += F * dy;
      }
    }
    a.vx = (a.vx + fx) * 0.5;
    a.vy = (a.vy + fy) * 0.5;
    a.x += a.vx;
    a.y += a.vy;
    if (a.x <= 0 || a.x >= width) { a.vx *= -1; }
    if (a.y <= 0 || a.y >= height) { a.vy *= -1; }
  }
};


yellow = create(400, "yellow");
red = create(400, "red");
green = create(400, "green");


const update = () => {
  rule(green, green, -0.32);
  rule(green, red, -0.17);
  rule(green, yellow, 0.34);
  rule(red, red, -0.1);
  rule(red, green, -0.34);
  rule(yellow, yellow, 0.15);
  rule(yellow, green, -0.2);
  m.clearRect(0, 0, width, height);
  draw(0, 0, "black", width, height);
  for (i = 0; i < atoms.length; i++) {
    draw(atoms[i].x, atoms[i].y, atoms[i].color, 10, 10);
  }
  requestAnimationFrame(update);
};


update();