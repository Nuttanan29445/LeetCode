/**
 * @param {number} x
 * @param {number} y
 * @return {string}
 */
var losingPlayer = function (x, y) {
  let win = 0;
  while (x >= 1 && y >= 4) {
    x -= 1;
    y -= 4;
    win += 1;
  }

  if (win % 2 === 0) {
    return "Bob";
  } else {
    return "Alice";
  }
};
