export const maxProfitMinMax = (prices: number[]): number => {
  let profit = 0;

  let localMin = prices[0];
  let localMax = prices[0];

  for (let i = 0; i < prices.length; i++) {
    if (prices[i] > localMax) {
      localMax = prices[i];
    } else {
      profit += localMax - localMin;
      localMin = prices[i];
      localMax = prices[i];
    }
  }

  profit += localMax - localMin;

  return profit;
};

export const maxProfitTrends = (prices: number[]): number => {
  let profit = 0;

  for (let i = 0; i < prices.length; i++) {
    if (prices[i] < prices[i + 1]) {
      profit += prices[i + 1] - prices[i];
    }
  }

  return profit;
};

console.log(maxProfitMinMax([7,1,5,3,6,4]))
