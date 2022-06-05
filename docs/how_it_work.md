# The Memo Format
The Memos are stored in markdown files and the given text input can also contain
mark down format tags.

## Input
```markdown
I am a **stupid** memo #target1
```

## Storage
Somewhere in a file target1.md
```markdown
* I am a **stupid** memo one [20220101T160627]
* I am a **stupid** memo two [20220201T160627]
* ~~and another memo, but this one was marked as ***done***~~ [20220301T160627] [Done: 20220401T100000]
```
