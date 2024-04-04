SELECT d.id AS draft_id,
       d.challengeID,
       d.userID1,
       d.userID2,
       d.status AS draft_status,
       dr.id AS round_id,
       dr.roundNumber,
       dr.status AS round_status,
       dr.winnerUserID,
       dd.deck
FROM drafts d
         JOIN draft_rounds dr ON d.id = dr.draftID
         JOIN draft_deck dd ON dr.id = dd.roundID
WHERE d.userID1 = {{.UserID}} OR d.userID2 = {{.UserID}};