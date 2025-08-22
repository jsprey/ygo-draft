import React, {useState} from "react";
import classNames from "classnames";
import YgoIcon from "../../../core/YgoIcon";
import {DraftMode} from "../../../api/Draft";
import Alert from "../../../core/Alert";

type SettingsModeSelectEntryProps = {
    value: DraftMode
    setValue: React.Dispatch<React.SetStateAction<DraftMode>>
    className?: string
    md: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
};

function SettingsModeSelectEntry(props: SettingsModeSelectEntryProps) {
    const [expanded, setExpanded] = useState<boolean>(false)

    const CollapsedIcon = <YgoIcon icon={"help-outline"}
                                   size={18}
                                   onClick={(event) => {
                                       setExpanded(true)
                                       event.preventDefault()
                                   }}
                                   classNames={classNames("fill-dark dark:fill-light", "hover:fill-primary", "active:hover:fill-primary-hover")}/>
    const ExpandIcon = <YgoIcon icon={"help-fill"}
                                size={18}
                                onClick={(event) => {
                                    setExpanded(false)
                                    event.preventDefault()
                                }}
                                classNames={classNames("fill-primary", "hover:fill-primary-hover", "active:hover:fill-primary-active")}/>

    const modeButtonsCN = classNames("p-2 border-t-2 border-b-2 border-border border-border dark:text-light")
    const modeButtonsSelectedCN = classNames("bg-primary text-light")
    return <div className={classNames(props.className ? props.className : "")}>
        <div className={"w-full"}>
            <div className={"flex mb-2"}>
                <div className={"self-center font-bold mr-2 dark:text-light"}>Draft Mode</div>
                {expanded ? ExpandIcon : CollapsedIcon}
            </div>
            {expanded ? <div>
                <Alert variant={"info"} className={"mb-2"}>Defines the mode for the draft:<br/>
                    <span className={"font-bold"}>Rounds</span> - Play a fixed amount of rounds.<br/>
                    <span className={"font-bold"}>BestOf</span> - Play a game until one player wins at leat half of the
                    rounds.</Alert>
            </div> : <></>}
        </div>
        <div>
            <button className={classNames("border-l-2 border-r-2 rounded-tl rounded-bl", modeButtonsCN, props.value === "rounds" ? modeButtonsSelectedCN : "")} onClick={event => {
                props.setValue("rounds")
                event.preventDefault()
            }
            }>
                Rounds
            </button>
            <button className={classNames("border-r-2 rounded-tr rounded-br", modeButtonsCN, props.value === "bestof" ? modeButtonsSelectedCN : "")} onClick={event => {
                props.setValue("bestof")
                event.preventDefault()
            }
            }>
                BestOf
            </button>
        </div>
    </div>
}

export default SettingsModeSelectEntry
