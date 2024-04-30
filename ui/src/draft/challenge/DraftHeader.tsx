import React, {useState} from "react";
import {Draft} from "../../api/Draft";
import DraftSettingsDetails from "./DraftSettingsDetails";
import classNames from "classnames";
import DraftBattleBadge from "./DraftBattleBadge";
import {User} from "../../api/hooks/users/useUsers";
import {Friend} from "../../api/UserModel";
import YgoIcon from "../../core/YgoIcon";

const CollapsedIcon = <YgoIcon icon={"collapse"} size={18}
                               classNames={"fill-dark dark:fill-light hover:fill-dark-1 hover:dark:fill-light-1 active:fill-dark-2 active:dark:fill-light-2"}/>
const ExpandIcon = <YgoIcon icon={"expand"} size={18}
                            classNames={"fill-dark dark:fill-light hover:fill-dark-1 hover:dark:fill-light-1 active:fill-dark-2 active:dark:fill-light-2"}/>

type DraftHeaderProps = {
    draft: Draft;
    player: User,
    enemy: Friend
};

function DraftHeader(props: DraftHeaderProps) {
    const [settingCollapsed, setSettingCollapsed] = useState<boolean>(true)

    const settingsHeaderCN = classNames("flex-grow pl-2", "rounded-tl rounded-tr", settingCollapsed ? "rounded-bl rounded-br" : "", "border border-light-3 dark:border-dark-3", "bg-gray-200 dark:bg-gray-600")
    const settingsBodyCN = classNames("p-2", "rounded-bl rounded-br", "border-l border-b border-r  border-light-3 dark:border-dark-3", "bg-lightI dark:bg-dark")

    return <div className={"w-full text-dark dark:text-light"}>
        <DraftBattleBadge playerNameOne={props.player.display_name} playerNameTwo={props.enemy.name}/>

        <div className={classNames("flex justify-center", "mt-2", settingsHeaderCN)}>
            <div
                className={classNames("self-center mr-1 uppercase fw-bold")}>
                Draft Settings
            </div>
            <div className={"self-center"} onClick={() => setSettingCollapsed(!settingCollapsed)}>
                {settingCollapsed ? CollapsedIcon : ExpandIcon}
            </div>
        </div>
        {settingCollapsed ? <></> :
            <DraftSettingsDetails settings={props.draft.settings} containerClasses={settingsBodyCN}/>}
    </div>
}

export default DraftHeader
