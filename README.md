# go-swap

Swap between binance and kucoin. Using trading strategy like
[Arbitrage Between Different Markets](https://github.com/Kucoin-academy/inter-exchange-spread/blob/master/README_EN.md):
"By purchasing the underlying asset in Exchange A and selling
it in Exchange B, traders can profit from price spreads."

## Binance and Kucoin

The project is using [binance](https://binance.com)
and [kucoin](https://www.kucoin.com)
exchanges to create orders to buy and sell.

### Transfer

The inter exchange spread strategy is commonly known as "arbitrage".

When there is a price difference between the two exchanges for the same
target, buy from Exchange A and sell on Exchange B to earn the price spread

The natural way to spin the story is to say that SKII(BTC)
in Korea(KuCoin Exchange) is 1000 yuan, while domestic
price(Huobi Exchange) is 1,800 yuan, and my freight
cost(transaction fee) is 200 yuan.
Then, I bought 10 SKII in Korea and sold them all in China.
Excluding the cost, I can earn a total price spread (1800-1000) * 10-200 = 7800 yuan.

[Read more](https://github.com/Kucoin-academy/inter-exchange-spread/blob/master/README_EN.md)

# TODO's

- [ ] Create order (buy/sell) in kucoin exchange;
- [ ] Create order (buy/sell) in binance exchange;
- [ ] Transfer currency kucoin -> binance;
- [ ] Transfer currency binance -> kucoin;
- [ ] Slack notification when happen any operation.

## 📫 Contributing

To contribute to the project, follow these steps:

1. Clone the repository: `git clone git@github.com:rafaelbmateus/go-binance-bot.git`
2. Create a feature branch: `git switch -c feature-a`
3. Make changes and confirm (try using [conventional commits](https://www.conventionalcommits.org)): `git commit -m 'feat: new feature'`
4. Push the feature branch: `git push origin feature-a`
5. Create a [pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
6. Get reviews from other users
7. Merge to `main` branch (we encourage using commit squash)
8. Remove the branch merged.
