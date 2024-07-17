/**
 * @description: 振动类型
 * @return {VibrateType}
 */
export enum VibrateType {
    EFFECT_TICK = 0 ,// 微妙触觉反馈的短暂振动(振动时长ms/振幅 (50ms/50))
    EFFECT_CLICK  = 1, // 中等振动，用于指示按下按钮或菜单选择(振动时长/振幅(75ms/100))
    EFFECT_HEAVY_CLICK  = 2, // 用于指示重要事件的长振动，例如通知或警报 (振动时长/振幅(250ms/255))
    EFFECT_DOUBLE_CLICK = 3 // 持续时间和幅度是 EFFECT_CLICK 两倍的振动，用于指示双击手势 (振动时长/振幅(150ms/250))
}

