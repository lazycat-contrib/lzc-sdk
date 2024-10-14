/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/pages/DeviceContacts.tsx
 */
import React from 'react'
import { BasePageProps } from '../types';

// import { ClientAuthorizationType } from '@lazycatcloud/sdk/extentions/client_authorization';
import { AppCommon } from '@lazycatcloud/sdk/extentions/app_common';

export type DeviceContactsProps = BasePageProps & {}

export default function DeviceContacts(props: DeviceContactsProps) {
    const { navigate, gateway: lzcAPI } = props

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
                <p>测试当前设备联系人相关接口</p>
            </div>

            <div style={{ display: 'inline-grid', justifyContent: 'center' }}>
                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            const currentDevice = await lzcAPI.currentDevice
                            const device = currentDevice!;
                            let reply = await device.contact.ListContacts({});
                            console.log('reply', reply);
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    ListContacts
                </button>

                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            const currentDevice = await lzcAPI.currentDevice
                            const device = currentDevice!;
                            let reply = await device.contact.AddContacts({
                                contacts: [
                                    {
                                        name: 'tzbin001',
                                        phones: ['10086']
                                    },
                                    {
                                        name: 'tzbin002',
                                        phones: ['1008611']
                                    }
                                ]
                            });
                            console.log('reply', reply);
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    AddContacts
                </button>

                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            const currentDevice = await lzcAPI.currentDevice
                            const device = currentDevice!;
                            const { contacts } = await device.contact.ListContacts({});
                            let newContacts = contacts.map((i) => {
                                if (i.name.includes('tzbin')) {
                                    return {
                                        ...i,
                                        phones: ['1000011']
                                    }
                                }
                            }).filter(i => i) as any[]
                            console.log('newContacts', newContacts);
                            let reply = await device.contact.UpdateContacts({
                                contacts: newContacts
                            });
                            console.log('reply', reply);
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    UpdateContact
                </button>

                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        try {
                            const currentDevice = await lzcAPI.currentDevice
                            const device = currentDevice!;
                            const { contacts } = await device.contact.ListContacts({});
                            let ids = contacts.map((i) => {
                                if (i.name.includes('tzbin')) {
                                    return i.id
                                }
                            }).filter(i => i) as string[]
                            console.log('ids', ids);
                            let reply = await device.contact.DeleteContacts({
                                ids: ids
                            });
                            console.log('reply', reply);
                        } catch (error) {
                            // throw error;
                            console.error('ListContacts error', error);
                        }
                    }}
                >
                    DeleteContacts
                </button>
            </div>

            <div>
                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        // const clientAuthorizationBaseStatusPromise = AppCommon.GetClientAuthorizationStatus(ClientAuthorizationType.Contacts);
                        // clientAuthorizationBaseStatusPromise.then(res => {
                        //     console.log(res)
                        // })
                    }}
                >
                    获取联系人状态
                </button>
            </div>
            <div>
                <button
                    style={{ marginTop: 10 }}
                    onClick={async () => {
                        // await AppCommon.RequestClientAuthorization(ClientAuthorizationType.Contacts)
                    }}
                >
                    申请联系人状态
                </button>
            </div>
        </div>
    )
}
