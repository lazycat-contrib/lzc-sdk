/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/pages/QueryAssetMatePage.tsx
 */
import React, { useState } from 'react'
import { EndDeviceProxy } from '@lazycatcloud/sdk/index';
import { AssetMetaReview, type PhotoMeta } from '../components';
import { BasePageProps } from '../types';

export default function QueryAssetMatePage(props: BasePageProps & {
    device: EndDeviceProxy
}) {
    const { navigate, device } = props;

    const [assetInfo, setAssetInfo] = useState<{
        id: string,
        message: string,
        mate?: PhotoMeta
    }>({
        id: '',
        message: '',
        mate: undefined,
    })

    function putMessage(msg: any) {
        setAssetInfo({
            ...assetInfo,
            message: `${msg}`
        })
    }

    return <div style={{ width: '100%' }}>
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
            <p>通过媒体 ID 查询图片 Mate 信息</p>
        </div>

        <div style={{ display: 'flex' }}>
            <input value={assetInfo.id} style={{ flex: 1, background: '#F1F1F1', border: "none", borderRadius: 5 }}
                onChange={(e) => {
                    const { target } = e;
                    if (target.value) {
                        setAssetInfo({
                            ...assetInfo,
                            id: target.value
                        })
                    } else {
                        setAssetInfo({
                            ...assetInfo,
                            id: ''
                        })
                    }
                }}
            ></input>
            <button style={{ marginLeft: 10 }} onClick={async () => {
                if (!assetInfo.id || assetInfo.id == '') {
                    putMessage(`[错误] 媒体 ID 为空`)
                    return
                }

                console.log('id', assetInfo.id, device);
                try {
                    const photo = await device.photolibrary.QueryPhoto({ id: assetInfo.id })
                    setAssetInfo({
                        ...assetInfo,
                        mate: photo,
                        message: JSON.stringify(photo, null, 2),
                    })
                } catch (error) {
                    putMessage(error)
                }

            }}>查询
            </button>
        </div>

        <div style={{ marginTop: 10 }}>

            <p>其他测试项</p>
            <button
                onClick={async () => {
                    console.log('currentDevice', device)
                    try {
                        let urls = await device.photolibrary.QueryAssetUrlPath({});
                        console.log('debug image urls', urls);
                        putMessage(`[QueryAssetUrlPath] ${JSON.stringify(urls, null, 2)}`)
                    } catch (error) {
                        putMessage(`[QueryAssetUrlPath 错误] ${error}`)
                        throw error;
                    }
                }}
            >
                QueryAssetUrlPath
            </button>
        </div>

        <div style={{ display: 'flex', flexDirection: 'column', marginTop: 25 }}>
            {assetInfo.message && (
                <>
                    <p style={{ textAlign: 'start' }}>查询结果</p>
                    <textarea disabled value={assetInfo.message} style={{ marginBottom: 15, height: 150 }}></textarea>
                </>
            )}
            {assetInfo.mate && (
                <AssetMetaReview meta={assetInfo.mate} />
            )}
        </div>
    </div>
}
