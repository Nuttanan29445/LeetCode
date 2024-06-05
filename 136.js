/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function (nums) {
  let d = {};
  for (let i of nums) {
    if (!d[i]) {
      d[i] = 0;
    }
    d[i] += 1;
  }
  const set = new Set(nums);
  for (let j of set) {
    if (d[j] === 1) {
      return j;
    }
  }
};
console.log(singleNumber([2, 2, 1]));
