/**
 * @param {number} start
 * @param {number} goal
 * @return {number}
 */
var minBitFlips = function (start, goal) {
  let binS = start.toString(2);
  let binG = goal.toString(2);
  const maxLen = Math.max(binS.length, binG.length);
  let s = "0".repeat(maxLen - binS.length) + binS;
  let g = "0".repeat(maxLen - binG.length) + binG;
  let ans = 0;
  for (let i = 0; i < maxLen; i++) {
    if (s[i] != g[i]) {
      ans += 1;
    }
  }
  return ans;
};
