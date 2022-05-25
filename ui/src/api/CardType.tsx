import {Card} from "./CardModel";

export enum CardType {
    Unknown,
    NormalMonster,
    NormalTunerMonster,
    EffectMonster,
    FlipEffectMonster,
    FlipTunerEffectMonster,
    GeminiMonster,
    PendulumEffectMonster,
    PendulumFlipEffectMonster,
    PendulumNormalMonster,
    PendulumTunerEffectMonster,
    RitualEffectMonster,
    RitualMonster,
    ToonMonster,
    SpiritMonster,
    TunerMonster,
    UnionEffectMonster,
    SkillCard,
    SpellCard,
    TrapCard,
    FusionMonster = 100,
    LinkMonster,
    PendulumEffectFusionMonster,
    SynchroMonster,
    SynchroPendulumEffectMonster,
    SynchroTunerMonster,
    XYZMonster,
    XYZPendulumEffectMonster,
}

export function GetCardType(card: Card): CardType {
    switch (card.type) {
        case "Normal Monster":
            return CardType.NormalMonster
        case "Effect Monster":
            return CardType.EffectMonster
        case "Flip Effect Monster":
            return CardType.FlipEffectMonster
        case "Flip Tuner Effect Monster":
            return CardType.FlipTunerEffectMonster
        case "Gemini Monster":
            return CardType.GeminiMonster
        case "Normal Tuner Monster":
            return CardType.NormalTunerMonster
        case "Pendulum Effect Monster":
            return CardType.PendulumEffectMonster
        case "Pendulum Flip Effect Monster":
            return CardType.PendulumFlipEffectMonster
        case "Pendulum Normal Monster":
            return CardType.PendulumNormalMonster
        case "Pendulum Tuner Effect Monster":
            return CardType.PendulumTunerEffectMonster
        case "Ritual Effect Monster":
            return CardType.RitualEffectMonster
        case "Ritual Monster":
            return CardType.RitualMonster
        case "Skill Card":
            return CardType.SkillCard
        case "Spell Card":
            return CardType.SpellCard
        case "Spirit Monster":
            return CardType.SpiritMonster
        case "Trap Card":
            return CardType.TrapCard
        case "Toon Monster":
            return CardType.ToonMonster
        case "Tuner Monster":
            return CardType.TunerMonster
        case "Union Effect Monster":
            return CardType.UnionEffectMonster
        case "Fusion Monster":
            return CardType.FusionMonster
        case "Link Monster":
            return CardType.LinkMonster
        case "Pendulum Effect Fusion Monster":
            return CardType.PendulumEffectFusionMonster
        case "Synchro Monster":
            return CardType.SynchroMonster
        case "Synchro Pendulum Effect Monster":
            return CardType.SynchroPendulumEffectMonster
        case "Synchro Tuner Monster":
            return CardType.SynchroTunerMonster
        case "XYZ Monster":
            return CardType.XYZMonster
        case "XYZ Pendulum Effect Monster":
            return CardType.XYZPendulumEffectMonster
        default:
            return CardType.Unknown
    }
}

export function GetCardTypeString(type: CardType): string {
    switch (type) {
        case CardType.NormalMonster:
            return "Normal Monster"
        case CardType.EffectMonster:
            return "Effect Monster"
        case CardType.FlipEffectMonster:
            return "Flip Effect Monster"
        case CardType.FlipTunerEffectMonster:
            return "Flip Tuner Effect Monster"
        case CardType.GeminiMonster:
            return "Gemini Monster"
        case CardType.NormalTunerMonster:
            return "Normal Tuner Monster"
        case CardType.PendulumEffectMonster:
            return "Pendulum Effect Monster"
        case CardType.PendulumFlipEffectMonster:
            return "Pendulum Flip Effect Monster"
        case CardType.PendulumNormalMonster:
            return "Pendulum Flip Effect Monster"
        case CardType.PendulumTunerEffectMonster:
            return "Pendulum Tuner Effect Monster"
        case CardType.RitualEffectMonster:
            return "Ritual Effect Monster"
        case CardType.RitualMonster:
            return "Ritual Monster"
        case CardType.SkillCard:
            return "Skill Card"
        case CardType.SpellCard:
            return "Spell Card"
        case CardType.SpiritMonster:
            return "Spirit Monster"
        case CardType.TrapCard:
            return "Trap Card"
        case CardType.ToonMonster:
            return "Toon Monster"
        case CardType.TunerMonster:
            return "Tuner Monster"
        case CardType.UnionEffectMonster:
            return "Union Effect Monster"
        case CardType.FusionMonster:
            return "Fusion Monster"
        case CardType.LinkMonster:
            return "Link Monster"
        case CardType.PendulumEffectFusionMonster:
            return "Pendulum Effect Fusion Monster"
        case CardType.SynchroMonster:
            return "Synchro Monster"
        case CardType.SynchroPendulumEffectMonster:
            return "Synchro Pendulum Effect Monster"
        case CardType.SynchroTunerMonster:
            return "Synchro Tuner Monster"
        case CardType.XYZMonster:
            return "XYZ Monster"
        case CardType.XYZPendulumEffectMonster:
            return "XYZ Pendulum Effect Monster"
        default:
            return "Unknown"
    }
}