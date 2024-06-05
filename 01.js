/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  for (let j = 0; j < nums.length; j++) {
    for (let k = j + 1; k < nums.length; k++) {
      if (parseInt(nums[j]) + parseInt(nums[k]) == target) {
        return [j, k];
      }
    }
  }
  return;
};

console.log(twoSum([2, 7, 11, 15], 9));
