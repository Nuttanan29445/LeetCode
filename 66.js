/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  let map = digits.map((digit) => String(digit));
  let ans = BigInt(map.join("")) + 1n;
  let z = [];
  for (let i of String(ans)) {
    z.push(i);
  }
  return z;
};

console.log(plusOne([6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3]));
