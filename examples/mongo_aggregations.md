# Mongo aggregations

## Find users and groups

```js
[
  {
    $match: {
      name: "John",
      surname: "Doe",
    },
  },
  {
    $lookup: {
      from: "groups",
      localField: "groupRefs",
      foreignField: "_id",
      as: "groups",
    },
  },
  {
    $unwind: "$groups",
  },
  {
    $group: {
      _id: "$_id",
      user: {
        $first: "$$ROOT",
      },
      groups: {
        $push: "$groups",
      },
    },
  },
  {
    $replaceRoot: {
      newRoot: {
        $mergeObjects: [
          "$user",
          {
            groups: "$groups",
          },
        ],
      },
    },
  },
  {
    $project: {
      groupRefs: 0,
    },
  },
];
```

## Find users and groups with group permissions

```js
[
  {
    $match: {
      name: "John",
      surname: "Doe",
    },
  },
  {
    $lookup: {
      from: "groups",
      localField: "groupRefs",
      foreignField: "_id",
      as: "groups",
    },
  },
  {
    $unwind: "$groups",
  },
  {
    $lookup: {
      from: "permissions",
      localField: "groups.permissionRefs",
      foreignField: "_id",
      as: "groups.permissions",
    },
  },
  {
    $project: {
      "groups.userRefs": 0,
      "groups.permissionRefs": 0,
    },
  },
  {
    $group: {
      _id: "$_id",
      user: {
        $first: "$$ROOT",
      },
      groups: {
        $push: "$groups",
      },
    },
  },
  {
    $replaceRoot: {
      newRoot: {
        $mergeObjects: [
          "$user",
          {
            groups: "$groups",
          },
        ],
      },
    },
  },
  {
    $project: {
      user: 0,
      groupRefs: 0,
    },
  },
];
```

## Find users who have not logged in last 6 months with their groups and group permissions

```js
[
  {
    $addFields: {
      lastSuccessfulLoginDate: {
        $toDate: "$userloginmetadata.lastsuccessfullogin",
      },
    },
  },
  {
    $match: {
      $expr: {
        $lt: [
          "$lastSuccessfulLoginDate",
          {
            $dateSubtract: {
              startDate: new Date(),
              unit: "month",
              amount: 6,
            },
          },
        ],
      },
    },
  },
  {
    $lookup: {
      from: "groups",
      localField: "groupRefs",
      foreignField: "_id",
      as: "groups",
    },
  },
  {
    $unwind: "$groups",
  },
  {
    $lookup: {
      from: "permissions",
      localField: "groups.permissionRefs",
      foreignField: "_id",
      as: "groups.permissions",
    },
  },
  {
    $project: {
      "groups.userRefs": 0,
      "groups.permissionRefs": 0,
    },
  },
  {
    $group: {
      _id: "$_id",
      user: {
        $first: "$$ROOT",
      },
      groups: {
        $push: "$groups",
      },
    },
  },
  {
    $replaceRoot: {
      newRoot: {
        $mergeObjects: [
          "$user",
          {
            groups: "$groups",
          },
        ],
      },
    },
  },
  {
    $project: {
      user: 0,
      groupRefs: 0,
    },
  },
];
```

## Find user and his coworkers

```js
[
  {
    $match: {
      name: "John",
      surname: "Doe",
    },
  },
  {
    $lookup: {
      from: "groups",
      localField: "groupRefs",
      foreignField: "_id",
      as: "groups",
    },
  },
  {
    $unwind: "$groups",
  },
  {
    $lookup: {
      from: "permissions",
      localField: "groups.permissionRefs",
      foreignField: "_id",
      as: "groups.permissions",
    },
  },
  {
    $lookup: {
      from: "users",
      localField: "groups.userRefs",
      foreignField: "_id",
      as: "groups.users",
    },
  },
  {
    $project: {
      "groups.userRefs": 0,
      "groups.permissionRefs": 0,
    },
  },
  {
    $group: {
      _id: "$_id",
      user: {
        $first: "$$ROOT",
      },
      groups: {
        $push: "$groups",
      },
    },
  },
  {
    $replaceRoot: {
      newRoot: {
        $mergeObjects: [
          "$user",
          {
            groups: "$groups",
          },
        ],
      },
    },
  },
  {
    $project: {
      user: 0,
      groupRefs: 0,
    },
  },
];
```
