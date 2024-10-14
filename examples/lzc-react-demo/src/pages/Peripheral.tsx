/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/pages/Peripheral.tsx
 */
import React, { useState } from 'react'
import { BasePageProps } from '../types';

import { FsType, Filesystem } from '@lazycatcloud/sdk/common/peripheral_device';

export type PeripheralProps = BasePageProps & {

}

export default function Peripheral(props: PeripheralProps) {
    const { navigate, gateway: lzcAPI } = props;

    const [fsList, setFsList] = useState<(Filesystem & { type: 'mounted' | 'unmounted' })[]>([])

    const remotes = {
        webdev: {
            address: 'http://127.0.0.1:8061',
        },
        smb: {
            address: '//127.0.0.1:445/sambshare',
            user: '',
            password: ''
        },
    }

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
                            lzcAPI.pd.MountWebDav({
                                mountpoint: "webdav1",
                                url: remotes.webdev.address,
                                username: '',
                                password: '',
                            })
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    Mount Remote WebDav
                </button>

                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            // const currentDevice = await lzcAPI.currentDevice
                            // const device = currentDevice!;
                            const call = await lzcAPI.pd.MountRemoteDisk({
                                mountpoint: 'smb1',
                                fstype: FsType.SAMBA,
                                remoteTarget: remotes.smb.address,
                                username: remotes.smb.user,
                                password: remotes.smb.password
                            })
                            console.log('call', call);

                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    Mount Remote Disk
                </button>

                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            // const currentDevice = await lzcAPI.currentDevice
                            // const device = currentDevice!;
                            const fsAll: any[] = [];
                            const list = await lzcAPI.pd.ListFilesystems({})
                            for (const item of list.mounted) {
                                fsAll.push({
                                    type: 'mounted',
                                    ...item
                                })
                            }
                            for (const item of list.umounted) {
                                fsAll.push({
                                    type: 'unmounted',
                                    ...item
                                })
                            }
                            console.log('list', list, fsAll);
                            setFsList(fsAll);
                            // const fs = list.mounted.filter(i => i.source == remotes.smb.address)[0]
                            // console.log('fs', fs);
                            // if (!fs) throw Error('fs is null')

                            // const call = await lzcAPI.pd.UmountFilesystem({
                            //     mountpoint: "/data/media/com.lzc.react.demo/smb1",
                            // })
                            // console.log('call', call);
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    List FileSystems
                </button>

                <div>
                    {fsList.map((i) => i ? (
                        <p key={i.uuid}>
                            {`[${i.fstype}-${i.uuid}] ${i.mountpoint}`} <a
                                href='javascript:void(0)'
                                onClick={async () => {
                                    console.warn('Not implemented');
                                }}
                            >{i.type == 'mounted' ? '卸载' : '挂载'}</a>
                        </p>
                    ) : undefined)}
                </div>

            </div>
        </div>
    )
}
