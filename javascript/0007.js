/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
  let z = [];
  let temp = Math.pow(2, 31);
  for (let i of String(x)) {
    z.push(i);
  }
  if (x > 0) {
    const ans = parseInt(z.reverse().join(""));
    if (ans > temp) {
      return 0;
    } else {
      return ans;
    }
  } else {
    const ans = parseInt(z[0] + z.slice(1).reverse().join(""));
    if (ans > temp) {
      return 0;
    } else {
      return ans;
    }
  }
};
console.log(reverse(-123));
