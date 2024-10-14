/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/pages/Peripheral.tsx
 */
import React from 'react'

import { BasePageProps } from '../types';

export type PeripheralProps = BasePageProps & {

}

export default function Peripheral(props: PeripheralProps) {
    const { navigate, gateway: lzcAPI } = props;
    return (
        <div style={{ width: '100%' }}>
            <div style={{
                width: '100%',
                display: 'flex',
                flexDirection: 'row',
                marginBottom: 10
            }}>
                <button
                    style={{ marginRight: 10 }}
                    onClick={async () => {
                        navigate.setPageIndex(undefined); // 返回主页
                    }}
                >
                    返回
                </button>
                <p>测试盒子外围设备相关接口</p>
            </div>
            <div style={{ display: 'inline-grid', justifyContent: 'center' }}>
                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            // const currentDevice = await lzcAPI.currentDevice
                            // const device = currentDevice!;
                            lzcAPI.pd.MountWebDav({})
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    Mount WebDav
                </button>
            </div>
        </div>
    )
}
