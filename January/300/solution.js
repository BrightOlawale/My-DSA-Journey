/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function (nums) {
  if (!nums.length) {
    return 0;
  }
  let subsq = [];
  subsq.push(nums[0]);

  for (let i = 1; i < nums.length; i++) {
    const num = nums[i];

    if (num < subsq[0]) {
      subsq[0] = num;
    } else if (num > subsq[subsq.length - 1]) {
      subsq.push(num);
    } else {
      let left = 0;
      let right = subsq.length - 1;

      while (left < right) {
        const mid = Math.floor((left + right) / 2);

        if (subsq[mid] < num) {
          left = mid + 1;
        } else {
          right = mid;
        }
      }

      subsq[left] = num;
    }
  }
  return subsq.length;
};
