/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  let z = [];
  let set = new Set(nums);
  for (let i of set) {
    z.push(nums.filter((n) => n === i).length);
  }
  let j = z.sort((a, b) => b - a).slice(0, k);
  let ans = [];
  for (let l of j) {
    for (let m of set) {
      if (nums.filter((n) => n === m).length === l) {
        if (!ans.includes(m)) {
          ans.push(m);
          break;
        }
      }
    }
  }
  return ans;
};

console.log(topKFrequent([1, 2], 2));
