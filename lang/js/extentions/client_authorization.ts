/*
 * @Author: Bin
 * @Date: 2024-05-14
 * @FilePath: /lzc-sdk/lang/js/extentions/client_authorization.ts
 */

/**
 * @description: 客户端权限状态
 * @return {ClientAuthorizationBaseStatus}
 */
export enum ClientAuthorizationBaseStatus {
    Unknown = 0, // 未知的（或用户未决定）
    Denied = 1, // 拒绝授权
    Authorized = 2, // 批准
    Limited = 3 // 有限授权（例如: Android 中未完全允许文件访问）
}

/**
 * @description: 客户端授权类型
 * @return {ClientAuthorizationType}
 */
export enum ClientAuthorizationType {
    PhotoLibrary = 'photo_library' // 图片库权限
}
