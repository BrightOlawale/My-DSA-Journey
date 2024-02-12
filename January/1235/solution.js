/**
 * @param {number[]} startTime
 * @param {number[]} endTime
 * @param {number[]} profit
 * @return {number}
 */
var jobScheduling = function (startTime, endTime, profit) {
  let len = profit.length;

  // Sort jobs based on their end times
  const jobs = [];
  for (let i = 0; i < len; i++) {
    jobs.push({
      startTime: startTime[i],
      endTime: endTime[i],
      profit: profit[i],
    });
  }
  jobs.sort((a, b) => a.endTime - b.endTime);

  // dp array to store the maximum profit at each index
  const dp = new Array(len).fill(0);
  dp[0] = jobs[0].profit;

  // Loop through the sorted jobs
  for (let i = 1; i < len; i++) {
    // Find the latest non-overlapping job
    let latestNonOverlapIndex = -1;
    for (let j = i - 1; j >= 0; j--) {
      if (jobs[j].endTime <= jobs[i].startTime) {
        latestNonOverlapIndex = j;
        break;
      }
    }

    dp[i] = Math.max(
      dp[i - 1],
      (latestNonOverlapIndex !== -1 ? dp[latestNonOverlapIndex] : 0) +
        jobs[i].profit
    );
  }

  return dp[len - 1];
};
