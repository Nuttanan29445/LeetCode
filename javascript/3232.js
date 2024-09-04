/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canAliceWin = function (nums) {
  let sum1 = 0;
  let sum2 = 0;
  for (let i of nums) {
    if (i <= 9) {
      sum1 += i;
    } else if (i > 9) {
      sum2 += i;
    }
  }

  if (sum1 != sum2) {
    return true;
  } else {
    return false;
  }
};
