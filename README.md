# Go WarcraftLogs
A simple Golang wrapper for Warcraft Logs APIs.

## Usage
Initialize `WarcraftLogs` client with API Token:

```golang
api := warcraftlogs.New("cb63bb62fbadb166657d20927a2335ae")
reports := api.ReportsForGuild("GuildName", warcraft.Realm_Silvermoon, warcraft.Region_EU)

for _, report := range reports {
  // Use report information
}
```
