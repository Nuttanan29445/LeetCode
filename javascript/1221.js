/**
 * @param {string} s
 * @return {number}
 */
var balancedStringSplit = function (s) {
  let m = {
    R: 0,
    L: 0,
  };
  let ans = 0;
  for (val of s) {
    m[val] += 1;

    if (m["R"] == m["L"]) {
      ans += 1;
    }
  }
  return ans;
};
