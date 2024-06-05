/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  nums2.forEach((num) => {
    nums1.push(num);
  });
  nums1.sort((a, b) => a - b);
  if (nums1.length % 2 === 0) {
    return (nums1[nums1.length / 2 - 1] + nums1[nums1.length / 2]) / 2;
  } else {
    return nums1[(nums1.length - 1) / 2];
  }
};
