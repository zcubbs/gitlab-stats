# Gitlab Stats

This CLI allows for synthesized CSV extract of Gitlab data:

- [x] Export Projects for a given group & its sub-groups
- [x] Export User data, and AccessLevel for a given Project or Group


## Usage

```bash
~ gitlab-stats export group data -id <group_id> 
```

> Must export Env Variables: `GS_GITLAB_PRIVATE_TOKEN`and `GS_GITLAB_URL`. Optional: `GS_GITLAB_ROOT_GROUP_ID` if not fed to cmd
