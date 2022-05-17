import React from "react";
import {DraftStages} from "./DeckDraftWizard";
import {Button, Form} from "react-bootstrap";

export type PageSettingsProps = {
    setCurrentStage: React.Dispatch<React.SetStateAction<DraftStages>>
}

function PageSettings(props: PageSettingsProps) {

    return <>
        <div className={"p-3 mt-0 bg-blue-200 rounded-2 shadow-md mb-2"}>
            At this page it is possible to configure multiple aspects of the drafting phase. Keep in mind to use the
            same options as your dueling partner.
        </div>

        <Form>
            <Form.Label>Range</Form.Label>
            <Form.Range/>
        </Form>
        <Button onClick={() => props.setCurrentStage(DraftStages.DraftMain)}>Next</Button>
    </>
}


export default PageSettings
