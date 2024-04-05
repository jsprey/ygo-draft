import React from "react";
import {Alert, Spinner} from "react-bootstrap";
import {useParams} from "react-router";
import {useDraft} from "../../api/hooks/drafts/useDraft";

type DraftOverviewPageParams = {
    id: string;
};

function DraftOverviewPage() {
    let params = useParams<DraftOverviewPageParams>() as DraftOverviewPageParams;
    const draftRequest = useDraft(params.id)

    return <>
        <div className={"flex place-content-center dark:text-white"}>
            {draftRequest.isLoading ? <Spinner animation={"border"} size={"sm"}/> : <></>}
            {draftRequest.error ? <Alert variant={"danger"} className={"mb-0"}>Failed to load draft!</Alert> : <></>}
            {draftRequest.data ? "Draft: " + JSON.stringify(draftRequest.data) : <></>}
        </div>
    </>
}


export default DraftOverviewPage
