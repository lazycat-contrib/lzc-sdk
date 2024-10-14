/*
 * @Author: Bin
 * @Date: 2024-10-14
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/components/AssetMetaReview.tsx
 */
import React from 'react';
import { PhotoMeta } from '@lazycatcloud/sdk/localdevice/photo';

export { type PhotoMeta }
export default function AssetMetaReview(props: { meta: PhotoMeta, deletePhoto?: Function }) {
    const { meta: imageMeta, deletePhoto } = props;

    function decodeThumbnailUrl(urlStr: string) {
        try {
            const url = new URL(urlStr);
            if (!url.searchParams.has('size') || !(url.searchParams.has('width') || url.searchParams.has('height'))) {
                // 不存在缩略图参数，添加默认参数
                url.searchParams.set('width', '200');
                url.searchParams.set('height', '200');
            }
            return url.toString()
        } catch (error) {
            return urlStr
        }
    }

    return <>
        <div>
            查询图片 <span style={{ width: 100 }}>({imageMeta.id})</span>:
            <a href={imageMeta.photoUrl} target={'_blank'}>
                查看原图
            </a>
            {'  '}
            <a href={decodeThumbnailUrl(imageMeta.thumbnailUrl)} target={'_blank'}>
                缩略图
            </a>
            {'  '}

            {deletePhoto && (
                <>
                    <a
                        onClick={() => {
                            deletePhoto && deletePhoto([imageMeta.id]);
                        }}
                    >
                        删除图片
                    </a>{' '}
                </>
            )}

            <p>{imageMeta.fileName}</p>
            <div style={{ padding: 2, background: '#0001' }}>
                <img style={{ width: 150 }}
                    src={imageMeta.thumbnailUrl ? decodeThumbnailUrl(imageMeta.thumbnailUrl) : imageMeta.photoUrl}
                // src={imageMeta.photoUrl}
                />
            </div>
        </div>
    </>
}
