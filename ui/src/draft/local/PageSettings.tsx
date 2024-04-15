import React, {useState} from "react";
import {Button, Form, Row} from "react-bootstrap";
import SettingsEntry from "./SettingsEntry";
import CardSetSelector from "./CardSetSelector";
import {CardSet} from "../../api/Sets";
import {DraftMode, DraftSettings} from "../../api/Draft";
import classNames from "classnames";
import SettingsModeSelectEntry from "./SettingsModeSelectEntry";

export type PageSettingsProps = {
    onSettingsSubmit: (settings: DraftSettings) => void
    submitButtonName: string
    local?: boolean
}

function PageSettings(props: PageSettingsProps) {
    const [generalDraftCardSets, setGeneralDraftCardSets] = useState<CardSet[]>([])
    const [mode, setMode] = useState<DraftMode>("rounds")
    const [numberOfRounds, setNumberOfRounds] = useState(10)
    const [numberOfRoundsError, setNumberOfRoundsError] = useState("")
    const [mainDraftRound, setMainDraftRound] = useState(5)
    const [mainDraftRoundError, setMainDraftRoundError] = useState("")
    const [mainDraftSize, setMainDraftSize] = useState(3)
    const [mainDraftSizeError, setMainDraftSizeError] = useState("");
    const [extraDraftRound, setExtraDraftRound] = useState(5)
    const [extraDraftRoundError, setExtraDraftRoundError] = useState("")
    const [extraDraftSize, setExtraDraftSize] = useState(2)
    const [extraDraftSizeError, setExtraDraftSizeError] = useState("")
    const [validated, setValidated] = useState(false)

    const isLocalSettings = props.local ? props.local : false

    const handleSubmit = (event: any) => {
        event.preventDefault();
        event.stopPropagation();

        if (mode !== undefined && mainDraftRoundError === "" && mainDraftSizeError === "" && mainDraftRoundError === "" && extraDraftRoundError === "" && extraDraftSizeError === "") {
            setValidated(true)

            const draftSettings: DraftSettings = {
                mode: mode,
                mode_value: numberOfRounds,
                main_deck_size: mainDraftSize,
                main_deck_draws: mainDraftRound,
                extra_deck_draws: extraDraftRound,
                extra_deck_size: extraDraftSize,
                sets: generalDraftCardSets
            }

            props.onSettingsSubmit(draftSettings)
        }
    };

    function showDraftInformation() {
        return <div className={"p-3 mt-0 bg-blue-200 rounded-2 shadow-md mb-2"}>
            {isLocalSettings ? "At this page it is possible to configure multiple aspects of the drafting phase. Keep in " +
                "mind to use the same options as your dueling partner." : "Define the settings of your challenge!"
            }
        </div>
    }

    const headingCN = classNames("title mb-2 text-xl border-bottom fw-bold dark:text-white")
    return <>
        {showDraftInformation()}

        <Form className={"mt-3"} onSubmit={handleSubmit} noValidate validated={validated}>
            <div className={classNames(headingCN, "mt-0")}>General</div>
            <Row>
            <SettingsModeSelectEntry value={mode} setValue={setMode} className={"mb-2"} md={6}/>
            <SettingsEntry value={numberOfRounds} setValue={setNumberOfRounds}
                           error={numberOfRoundsError}
                           setError={setNumberOfRoundsError} min={2} max={20}
                           title={"Number of Rounds"}
                           className={"mb-2"}
                           tooltip={"Defines the number of bestof/normal rounds."}
                           md={6}/>
            </Row>

            <CardSetSelector selectedSets={generalDraftCardSets} setSelectedSets={setGeneralDraftCardSets}
                             rowClass={"mt-2"}
                             tooltip={"Only cards from the defined sets are used when drafting a deck."}></CardSetSelector>
            <div className={classNames(headingCN, "mt-4")}>Main Draft</div>
            <Row>
                <SettingsEntry value={mainDraftRound} setValue={setMainDraftRound} md={6}
                               error={mainDraftRoundError}
                               setError={setMainDraftRoundError} min={5} max={80}
                               title={"Number of Draws"}
                               tooltip={"Defines the number of draw rounds while drafting the main deck. The resulting main deck will have the same size as the draw number. Valid Values: [40-80]."}/>

                <SettingsEntry value={mainDraftSize} setValue={setMainDraftSize} md={6} error={mainDraftSizeError}
                               setError={setMainDraftSizeError} min={2} max={10}
                               title={"Card Each Draw"}
                               tooltip={"Defines the number of cards that are proposed for every draw round of the draft while drafting the main deck. Valid Values: [2-10]."}/>
            </Row>
            <div className={classNames(headingCN, "mt-4")}>Extra Draft</div>
            <Row>
                <SettingsEntry value={extraDraftRound} setValue={setExtraDraftRound} md={6}
                               error={extraDraftRoundError}
                               setError={setExtraDraftRoundError} min={0} max={20}
                               title={"Number of Draws"}
                               tooltip={"Defines the number of draw rounds while drafting the extra deck. The resulting extra deck will have the same size as the draw number. Valid Values: [0-20]."}/>

                <SettingsEntry value={extraDraftSize} setValue={setExtraDraftSize} md={6}
                               error={extraDraftSizeError}
                               setError={setExtraDraftSizeError} min={2} max={10}
                               title={"Card Each Draw"}
                               tooltip={"Defines the number of cards that are proposed for every draw of the draft while drafting the extra deck. Valid Values: [2-10]."}/>
            </Row>
            <div className={"flex place-content-end"}>
                <Button className={"mt-3"} type="submit"
                        disabled={mainDraftSizeError !== "" || mainDraftRoundError !== "" || extraDraftSizeError !== "" || extraDraftRoundError !== ""}>
                    {props.submitButtonName}
                </Button>
            </div>
        </Form>
    </>
}

export default PageSettings
