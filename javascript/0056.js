/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function (intervals) {
  intervals.sort((a, b) => a[0] - b[0]);

  const ans = [];
  let i = 0;

  while (i < intervals.length) {
    let start = i;
    let valStart = intervals[start][0];
    let valEnd = intervals[start][1];

    while (i + 1 < intervals.length && intervals[i + 1][0] <= valEnd) {
      valEnd = Math.max(valEnd, intervals[i + 1][1]);
      i++;
    }

    ans.push([valStart, valEnd]);
    i++;
  }

  return ans;
};
