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

var merge2 = function (intervals) {
  intervals.sort((a, b) => a[0] - b[0]);

  let ans = [intervals[0]];

  intervals.slice(1).forEach((val) => {
    let lastMerge = ans[ans.length - 1];

    if (val[0] <= lastMerge[1]) {
      lastMerge[1] = Math.max(lastMerge[1], val[1]);
    } else {
      ans.push(val);
    }
  });

  return ans;
};
