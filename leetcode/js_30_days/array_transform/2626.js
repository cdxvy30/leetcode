/**
 * @param {number[]} nums
 * @param {Function} fn
 * @param {number} init
 * @return {number}
 */
var reduce = function(nums, fn, init) {
  var ans = init;
  for (const num of nums) {
    ans = fn(ans, num);
  }
  return ans;
};