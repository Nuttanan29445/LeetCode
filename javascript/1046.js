/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function (stones) {
  while (stones.length > 1) {
    r;
    stones.sort((a, b) => b - a);

    let max1 = stones[0];
    let max2 = stones[1];

    if (max1 === max2) {
      stones = stones.slice(2);
    } else {
      stones[0] = max1 - max2;
      stones = stones.slice(0, 1).concat(stones.slice(2));
    }
  }

  return stones.length === 1 ? stones[0] : 0;
};
