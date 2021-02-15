# Usage

Mails daily price changes for a set watchlist.

```sh
bazel run 
STOCKBUDDY_PASSWORD=$STOCKBUDDY_PASSWORD bazel run --action_env=STOCKBUDDY_PASSWORD //summary:dailyprices -- --mail_to=<EMAIL_LIST>
```
