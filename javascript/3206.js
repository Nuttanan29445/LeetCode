/**
 * @param {number[]} colors
 * @return {number}
 */
var numberOfAlternatingGroups = function (colors) {
  let count = 0;
  for (let i = 0; i < colors.length; i++) {
    let index1;
    let index2;
    if (i + 2 >= colors.length) {
      index1 = (i + 2) % colors.length;
    } else {
      index1 = i + 2;
    }
    if (i + 1 >= colors.length) {
      index2 = (i + 1) % colors.length;
    } else {
      index2 = i + 1;
    }
    if (colors[i] === colors[index1] && colors[i] != colors[index2]) {
      count += 1;
    }
  }
  return count;
};
