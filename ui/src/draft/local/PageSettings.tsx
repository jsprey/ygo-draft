import React, {useState} from "react";
import SettingsEntry from "./SettingsEntry";
import CardSetSelector from "./CardSetSelector";
import {CardSet} from "../../api/Sets";
import {DraftMode, DraftSettings} from "../../api/Draft";
import classNames from "classnames";
import SettingsModeSelectEntry from "./SettingsModeSelectEntry";
import Alert from "../../core/Alert";
import Button from "../../core/Button";

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

    const isLocalSettings = props.local ? props.local : false

    function isValid(): boolean {
        return mode !== undefined && numberOfRoundsError === "" && mainDraftRoundError === "" && mainDraftSizeError === "" && extraDraftRoundError === "" && extraDraftSizeError === "";
    }

    const handleSubmit = () => {
        if (isValid()) {

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
        return isLocalSettings ? <Alert variant={"info"} className={"mt-2 mb-2"}>
            {isLocalSettings ? "At this page it is possible to configure multiple aspects of the drafting phase. Keep in " +
                "mind to use the same options as your dueling partner." : "Define the settings of your challenge!"
            }
        </Alert> : null
    }

    const headingCN = classNames("title mb-2 text-xl border-b border-border font-bold dark:text-white")
    return <>
        {showDraftInformation()}

        <div className={"mt-3"} onSubmit={handleSubmit}>
            <div className={classNames(headingCN, "mt-0")}>General</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsModeSelectEntry value={mode} setValue={setMode} className={"mb-2 pr-1"} md={6}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={numberOfRounds}
                                   setValue={setNumberOfRounds}
                                   error={numberOfRoundsError}
                                   setError={setNumberOfRoundsError}
                                   min={2} max={20}
                                   label={"Number of Rounds"}
                                   className={"mb-2 pl-1"}
                                   tooltip={"Defines the number of bestof/normal rounds."}/>
                </div>
            </div>

            <CardSetSelector selectedSets={generalDraftCardSets} setSelectedSets={setGeneralDraftCardSets}
                             rowClass={"mt-2 w-full"}
                             tooltip={"Only cards from the defined sets are used when drafting a deck."}/>
            <div className={classNames(headingCN, "mt-4")}>Main Draft</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsEntry value={mainDraftRound}
                                   setValue={setMainDraftRound}
                                   min={5} max={80}
                                   error={mainDraftRoundError}
                                   setError={setMainDraftRoundError}
                                   label={"Number of Draws"}
                                   className={"pr-1"}
                                   tooltip={"Defines the number of draw rounds while drafting the main deck. The resulting main deck will have the same size as the draw number. Valid Values: [40-80]."}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={mainDraftSize} setValue={setMainDraftSize}
                                   min={2} max={10}
                                   error={mainDraftSizeError}
                                   setError={setMainDraftSizeError}
                                   label={"Card Each Draw"}
                                   className={"pl-1"}
                                   tooltip={"Defines the number of cards that are proposed for every draw round of the draft while drafting the main deck. Valid Values: [2-10]."}/>
                </div>

            </div>
            <div className={classNames(headingCN, "mt-4")}>Extra Draft</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsEntry value={extraDraftRound}
                                   setValue={setExtraDraftRound}
                                   error={extraDraftRoundError}
                                   setError={setExtraDraftRoundError}
                                   min={0} max={20}
                                   className={"pr-1"}
                                   label={"Number of Draws"}
                                   tooltip={"Defines the number of draw rounds while drafting the extra deck. The resulting extra deck will have the same size as the draw number. Valid Values: [0-20]."}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={extraDraftSize}
                                   setValue={setExtraDraftSize}
                                   error={extraDraftSizeError}
                                   setError={setExtraDraftSizeError}
                                   min={2} max={10}
                                   className={"pl-1"}
                                   label={"Card Each Draw"}
                                   tooltip={"Defines the number of cards that are proposed for every draw of the draft while drafting the extra deck. Valid Values: [2-10]."}/>
                </div>
            </div>
            <div className={"flex place-content-end"}>
                <Button variant={"primary"} className={"mt-3"} onClick={handleSubmit}
                        disabled={!isValid()}>
                    {props.submitButtonName}
                </Button>
            </div>
        </div>
    </>
}

export default PageSettings
