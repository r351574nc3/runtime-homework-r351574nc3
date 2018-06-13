# Calculating support metrics

Customer support cases are escalated from our support team to a member of the Runtime Team when the case requires deeper technical investigation. The goal of this project is to calculate the hours we spend actively working on support cases. This will help us track our team's support burden over time.

Support cases are represented as a series of events. There are two types of events: either the assignee changed or the state changed. Support cases can be in four states: “null”, open, pending, or closed. All timestamps are UTC. 

**Feature Request**

Write a web service that can evaluate a set of support case events and produce a value for each support case that represents the amount of hours where both these statements are true: state==open and team==Runtime. 

For example, given the following events for support case 100, the correct output is 1+24=25 hours. There are two segments of time during which state==open and team==Runtime. The first segment (1 hour) begins when the case is assigned to Runtime and ends when the state is changed to pending. The second segment (24 hours) begins when the state is changed to open again and ends when the state is changed to closed.

```
[{"case_id": 100, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-01T00:00:00Z"},
 {"case_id": 100, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-01T00:00:00Z"},
 {"case_id": 100, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-02T15:00:00Z"},
 {"case_id": 100, "state": {"from": "open", "to": "pending"}, "timestamp": "2017-01-02T16:00:00Z"},
 {"case_id": 100, "state": {"from": "pending", "to": "open"}, "timestamp": "2017-01-07T06:00:00Z"},
 {"case_id": 100, "state": {"from": "open", "to": "closed"}, "timestamp": "2017-01-08T06:00:00Z"}]
```

**Acceptance Criteria**

Accept POST input of JSON support case events.  Return JSON output with the total number of hours each case was open and assigned to the Runtime Team.

Your web service should yield the following output when provided with the input.json file below.

```
curl -X POST http://localhost:8080/cases -H "Content-Type: application/json" -d @input.json
[{"case_id": 100, "hours": 25},
 {"case_id": 101, "hours": 0},
 {"case_id": 102, "hours": 268},
 {"case_id": 103, "hours": 104},
 {"case_id": 104, "hours": 45}]
```

input.json:

```
[
  {"case_id": 102, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-09T03:00:00Z"},
  {"case_id": 100, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-01T00:00:00Z"},
  {"case_id": 101, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-05T10:00:00Z"},
  {"case_id": 100, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-01T00:00:00Z"},
  {"case_id": 100, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-02T15:00:00Z"},
  {"case_id": 102, "state": {"from": "open", "to": "pending"}, "timestamp": "2017-01-09T11:00:00Z"},
  {"case_id": 100, "state": {"from": "open", "to": "pending"}, "timestamp": "2017-01-02T16:00:00Z"},
  {"case_id": 100, "state": {"from": "pending", "to": "open"}, "timestamp": "2017-01-07T06:00:00Z"},
  {"case_id": 104, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-13T06:00:00Z"},
  {"case_id": 100, "state": {"from": "open", "to": "closed"}, "timestamp": "2017-01-08T06:00:00Z"},
  {"case_id": 102, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-10T06:00:00Z"},
  {"case_id": 101, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-05T10:00:00Z"},
  {"case_id": 103, "state": {"from": null, "to": "open"}, "timestamp": "2017-01-11T19:00:00Z"},
  {"case_id": 102, "state": {"from": "pending", "to": "open"}, "timestamp": "2017-01-10T07:00:00Z"},
  {"case_id": 102, "state": {"from": "open", "to": "closed"}, "timestamp": "2017-01-21T11:00:00Z"},
  {"case_id": 103, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-11T22:00:00Z"},
  {"case_id": 101, "state": {"from": "open", "to": "closed"}, "timestamp": "2017-01-05T23:00:00Z"},
  {"case_id": 103, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-12T02:00:00Z"},
  {"case_id": 104, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-13T13:00:00Z"},
  {"case_id": 103, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-12T15:00:00Z"},
  {"case_id": 103, "state": {"from": "open", "to": "closed"}, "timestamp": "2017-01-16T19:00:00Z"},
  {"case_id": 104, "assignee": "Rabbit", "team": "Runtime", "timestamp": "2017-01-13T15:00:00Z"},
  {"case_id": 104, "assignee": "Badger", "team": "Runtime", "timestamp": "2017-01-15T08:00:00Z"},
  {"case_id": 104, "assignee": "Otter", "team": "Support", "timestamp": "2017-01-15T12:00:00Z"}
]
```
