import React, { useEffect, useRef, useState } from 'react';
import reactLogo from './assets/react.svg';
import './App.css';
import { LoadingIcon } from 'tdesign-icons-react';
// import { firstValueFrom } from 'rxjs';
import { EndDeviceProxy, lzcAPIGateway } from '@lazycatcloud/sdk/index';
import { ListAlbumsReply, ListAssetsSortType, PhotoMeta } from '@lazycatcloud/sdk/localdevice/photo';
import { Permission } from '@lazycatcloud/sdk/localdevice/permission';
import { AppCommon } from '@lazycatcloud/sdk/extentions/app_common';
// import {VibrateType} from "@lazycatcloud/sdk/extentions/vibrate_type";

import { AssetMetaReview } from './components';
import { PeripheralPage, DeviceContactsPage, QueryAssetMatePage } from './pages';

function App() {
    const [loading, setLoading] = useState(false);
    const [netState, setNetState] = useState('无网络');
    const lzcAPI = useRef<lzcAPIGateway>();
    const currentDevice = useRef<EndDeviceProxy>();

    const [imageMeta, setImageMeta] = useState<PhotoMeta>();
    const images = useRef<PhotoMeta[]>(Array<PhotoMeta>());
    const imageIndex = useRef(0);
    const [loadingImage, setLoadingImage] = useState(false);

    const [notImageUrl, setNotImageUrl] = useState(0);

    useEffect(() => {
        getOpenAppMethod()
    }, []);

    useEffect(() => {
        if (!lzcAPI.current) {
            lzcAPI.current = new lzcAPIGateway();
            // @ts-ignore
            window.ll = lzcAPI.current
            // 调试
            // getAlbumsDebug();

            if (!currentDevice.current) {
                getCurrentDevice()
            }
        }
    }, []);

    const netStateTransformStr = (state: number) => {
        switch (state) {
            case 1:
                return '无网络';
            case 2:
                return '以太网';
            case 3:
                return 'WIFI 网络';
            case 4:
                return '移动网络';
            case 5:
                return '2G 网络';
            case 6:
                return '3G 网络';
            case 7:
                return '4G 网络';
            case 8:
                return '5G 网络';
            default:
                return '未知';
        }
    };

    // 获取当前设备
    const getCurrentDevice: () => Promise<EndDeviceProxy> = async () => {
        if (!currentDevice.current) {
            const allPermissions = [
                Permission.CLIPBOARD,
                Permission.DEVICE_INFO,
                Permission.DOCUMENT,
                Permission.NETWORK_INFO,
                Permission.OPEN_DIALOG,
            ]


            for (let key in Permission) {
                if (typeof key === "number") {
                    var pType: any = Permission[key];
                    var pEnum: Permission = pType;
                    allPermissions.push(pEnum)
                }
            }

            console.log("allPermissions", allPermissions);


            currentDevice.current = (await GrantDevicePermission(allPermissions))!;
        }
        return currentDevice.current;
    };

    async function GrantDevicePermission(
        permissions: Permission | ArrayLike<Permission>,
    ): Promise<EndDeviceProxy> {
        let allRequiredPermissions: Permission[]
        if (Array.isArray(permissions)) {
            allRequiredPermissions = permissions
        } else {
            allRequiredPermissions = [permissions as Permission]
        }

        try {
            const gw = new lzcAPIGateway()
            const currentDevice = await gw.currentDevice

            for (let permission of allRequiredPermissions) {
                const hasPermission = await currentDevice.permission.GetPermission({
                    permission,
                })
                if (hasPermission.result) continue

                const success = await currentDevice.permission.RequestPermission({
                    permission,
                })

                if (!success.result) {
                    throw new Error("request local device permission failed")
                }
            }
            return currentDevice
        } catch (e) {
            throw e
        }
    }

    // 调试获取相册列表
    const getAlbumsDebug: () => Promise<ListAlbumsReply | undefined> = async () => {
        var albums: ListAlbumsReply | undefined;
        try {
            const device = await getCurrentDevice();
            albums = await device?.photolibrary.ListAlbums({});
            console.log('媒体列表', albums);
        } catch (error) {
            console.log('获取异常', error);
            throw error;
        }
        return albums;
    };

    // 测试获取相册中的资源
    const getAlbumDebug = async (ids: string[]) => {
        let album: PhotoMeta[] = [];
        try {
            const device = await getCurrentDevice();
            let photoMetas = device.photolibrary.ListPhotoMetas({
                albumIds: ids,
                thumbnailWidth: 100,
                thumbnailHeight: 100,
                needFileName: true,
                needAlbumIds: true,
            });

            await photoMetas.forEach(async (item) => {

                album.push(item);
                images.current.push(item);

                // 筛选视频
                // if (item.mime == "video/quicktime") {
                //     console.log("mime", item.mime);
                //     album.push(item);
                //     images.current.push(item);
                // }
            });

            return album;
        } catch (error) {
            throw error;
        }
    };

    // 测试获取相册中的资源
    const getImageDebug = async (id: string) => {
        try {
            const device = await getCurrentDevice();
            let photo = device.photolibrary.QueryPhoto({
                id,
            });
            return photo;
        } catch (error) {
            throw error;
        }
    };

    // 查询一张图片
    const getImageHash = async () => {
        try {
            const device = await getCurrentDevice();
            let photo = '';
            return photo;
        } catch (error) {
            throw error;
        }
    };

    // 删除图片
    const deletePhoto = async (ids: string[]) => {
        try {
            const device = await getCurrentDevice();
            const reply = await device.photolibrary.DeletePhoto({
                id: ids,
            });
            console.log('reply.failedId', reply);
        } catch (error) {
            console.log('失败', error);
        }
    };

    // 保存图片
    const putPhoto = async () => {
        try {
            const device = await getCurrentDevice();
            const { albums } = await device?.photolibrary.ListAlbums({});
            const album = albums[0];
            console.log('保存相册ID', album.title, album.id);
            if (!imageMeta) {
                if (images.current.length < 1) {
                    images.current = await getAlbumDebug([]);
                }
                const imageList = images.current;
                setImageMeta(imageList[imageIndex.current]);
            }

            device.photolibrary.PutPhoto({
                albumId: album.id,
                url: imageMeta?.photoUrl || imageMeta?.thumbnailUrl,
                fileName: 'test007',
            });
        } catch (error) {
            console.log('失败', error);
        }
    };

    // 获取文件打开方式
    const getFileOpenHandler = async () => {
        try {
            const device = await getCurrentDevice();
            const data = device.fileHandler.query({
                mime: 'text/plain',
            });
            console.log('文件打开方式', data);
        } catch (error) {
            console.log('失败', error);
        }
    };

    // useEffect(() => {
    // 	return () => {
    // 		URL.revokeObjectURL(imageMeta?.thumbnailUrl!);
    // 	};
    // }, []);

    const getOpenAppMethod = async () => {
        try {
            const device = await getCurrentDevice();

        } catch (error) {
            console.log("appWay error", error);
        }
    }

    const [pageIndex, setPageIndex] = useState<number | undefined>(undefined);
    const pageViews = [
        {
            view: (
                <>
                    <QueryAssetMatePage
                        navigate={{
                            setPageIndex,
                        }}
                        gateway={lzcAPI.current!}
                        device={currentDevice.current!}
                    />
                </>
            )
        },
        {
            view: (
                <DeviceContactsPage
                    navigate={{
                        setPageIndex,
                    }}
                    gateway={lzcAPI.current!}
                />
            )
        },
        {
            view: (
                <PeripheralPage
                    navigate={{
                        setPageIndex,
                    }}
                    gateway={lzcAPI.current!}
                />
            )
        }
    ];

    async function deletePhotoFor500() {
        try {
            if (images.current.length < 500) {
                console.log('deletePhotoFor500 failed', '数量不足');
                return
            }

            // 获取 500 个资源
            let list = images.current
            list = list.splice(0, 500);
            console.log('list', list);

            const ids = list.map((i) => i.id)
            const device = await getCurrentDevice();
            const reply = await device.photolibrary.DeletePhoto({
                id: ids,
            });
            console.log('reply.failedId', reply);

        } catch (error) {
            console.log('deletePhotoFor500 failed', error);
        }
    }

    return (
        <div className="App" style={{ width: '100%' }}>
            {pageIndex !== undefined && pageIndex < pageViews.length ? pageViews[pageIndex].view : (
                <>
                    <div>
                        <a href="https://vitejs.dev" target="_blank">
                            <img src="/vite.svg" className="logo" alt="Vite logo" />
                        </a>
                        <a href="https://reactjs.org" target="_blank">
                            <img src={reactLogo} className="logo react" alt="React logo" />
                        </a>
                    </div>
                    <h1>LzcAPI for React</h1>
                    <div className="card">
                        <div style={{ display: 'inline-grid', justifyContent: 'center' }}>

                            <p>--- box api tests ---</p>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    setPageIndex(2); // 打开测试盒子挂载外围设备接口页面
                                }}
                            >
                                Box Peripheral Tests
                            </button>

                            <p style={{ marginTop: 55 }}>--- local device api tests ---</p>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    setPageIndex(1); // 打开测试当前设备联系人接口页面
                                }}
                            >
                                Device Contacts Tests
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    const currentDevice = await lzcAPI.current?.currentDevice
                                    setPageIndex(0); // 打开查询媒体资源页面
                                }}
                            >
                                Device QueryAsset Tests
                            </button>

                            <p style={{ marginTop: 55 }}>--- other tests ---</p>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    try {
                                        const currentDevice = await lzcAPI.current?.currentDevice
                                        const device = currentDevice!;
                                        let reply = device.photolibrary.PutPhoto({
                                            url: 'https://open.doc.cxkjedu.com/logo.svg',
                                            // url: 'https://open.doc.cxkjedu.com/logo-with-shadow.png',
                                            fileName: 'test.svg'
                                        });
                                        console.log('reply', reply);
                                        await reply.forEach((item) => {
                                            console.log('reply', item);
                                        })
                                    } catch (error) {
                                        // throw error;
                                        console.error('PutPhoto error', error);
                                    }
                                }}
                            >
                                PutPhoto
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={deletePhotoFor500}
                            >
                                deletePhotoFor500
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    // let a = await AppCommon.ShareMedia({
                                    //     id: '4FD65BF9-A838-405C-A476-1F71CACBD587/L0/001',
                                    //     // ids: ['4FD65BF9-A838-405C-A476-1F71CACBD587/L0/001']
                                    // })
                                    // let a = await AppCommon.ShareWith("", "test", "")
                                    // let a = await AppCommon.ShareWith("", "/private/var/mobile/Containers/Shared/AppGroup/05277D27-8A5D-464E-8FD2-A496A4D1582C/Documents/懒猫网盘离线/0o0/admin/兔子2.0.zip", "")
                                    let ok = await AppCommon.ShareWithFiles(undefined, [
                                        "/private/var/mobile/Containers/Shared/AppGroup/05277D27-8A5D-464E-8FD2-A496A4D1582C/Documents/懒猫网盘离线/0o0/admin/IMG_0127.PNG",
                                        "/private/var/mobile/Containers/Shared/AppGroup/05277D27-8A5D-464E-8FD2-A496A4D1582C/Documents/懒猫网盘离线/0o0/admin/IMG_a0131.JPG"
                                    ])
                                    console.log('ShareWithFiles', ok);
                                }}
                            >
                                ShareMedia
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    const device = await lzcAPI.current?.currentDevice;
                                    if (!device) throw Error('device is null')
                                    device.fileHandler.copyFolder({
                                        "boxName": "lnks888",
                                        "devicePath": "/Users/mac/lzc-client-desktop",
                                        "username": "meetsong",
                                        "password": "",
                                        targetPath: "/backup",
                                        "webdavAddr": "https://file.lnks888.heiyu.space/dav/"
                                    }).forEach(task => {
                                        const on = JSON.stringify(task)
                                        if (task.msg) {
                                            console.log("收到同步信息错误:" + task.msg)
                                        } else {
                                            console.log("收到同步信息:" + on)
                                        }
                                    }).catch(res => {
                                        console.log("同步错误:" + res.toString())
                                    })
                                }}
                            >
                                SyncFolder
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    if (loading) return;
                                    let net = '';
                                    setLoading(true);
                                    try {
                                        const device = await lzcAPI.current?.currentDevice;
                                        const network = await device?.network.Query({}, undefined);
                                        net = netStateTransformStr(network?.ctype || -1);
                                    } catch (error) {
                                        net = '网络状态获取失败 ' + error;
                                        console.log('lzcAPI', '网络状态获取失败', error);
                                    }
                                    setLoading(false);
                                    setNetState(net);
                                }}
                            >
                                Get NetState {loading && <LoadingIcon size="1em" />}
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    getAlbumsDebug();
                                }}
                            >
                                Test Get Albums
                            </button>
                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    try {
                                        const devices = await lzcAPI.current?.devices.ListEndDevices({
                                            uid: (await lzcAPI.current.session).uid,
                                        });
                                        console.log('lzcAPI', '设备信息', devices);
                                    } catch (error) {
                                        console.log('lzcAPI', '设备信息', error);
                                    }
                                }}
                            >
                                Get Device List
                            </button>
                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    try {
                                        let { albums } = (await getAlbumsDebug())!;
                                        if (albums.length < 1) {
                                            throw Error('相册数量过少');
                                        }
                                        // let album = await getAlbumDebug([albums[albums.length - 1].id]);
                                        // let album = await getAlbumDebug(['7B874C5F-490B-42EE-911E-1CB55333056B/L0/040']);
                                        let album = await getAlbumDebug([]);
                                        images.current = album;
                                        console.log('获取相册', album);
                                    } catch (error) {
                                        console.log('获取失败', error);
                                    }
                                }}
                            >
                                Get All Album {`(${(images.current || []).length})`}
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                disabled={loadingImage}
                                onClick={async () => {
                                    if (loadingImage) {
                                        return;
                                    }

                                    try {
                                        setLoadingImage(true);
                                        if (images.current.length < 1) {
                                            getAlbumDebug([]);
                                        }
                                        let imageList = images.current;
                                        if (imageList.length < 1) {
                                            throw Error('相片数量过少');
                                        }
                                        let image = await getImageDebug(imageList[imageIndex.current].id);
                                        // let image = await getImageDebug('1FC00A27-FF03-48DF-9DBB-5F30DCB26FAD/L0/001'); // IMG_0009.MOV
                                        // let image = await getImageDebug('32354'); // IMG_0088.HEIC
                                        setImageMeta(image);
                                        imageIndex.current += 1;
                                        if (imageIndex.current > imageList.length - 1) {
                                            imageIndex.current = 0;
                                        }

                                        // let image = await getImageDebug('4D9CCB9F-0902-4572-95F2-BDB90F0A6E7F/L0/001');
                                        console.log('相片', imageIndex.current, image);
                                        // setImageUrl(image.photoUrl);
                                        setLoadingImage(false);
                                    } catch (error) {
                                        console.log('获取失败', error);
                                        setLoadingImage(false);
                                    }
                                }}
                            >
                                Get {imageIndex.current == 0 ? 'First' : 'Next'} Image{' '}
                                {loadingImage && <LoadingIcon size="1em" />}
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                disabled={loadingImage}
                                onClick={async () => {
                                    try {
                                        setLoadingImage(true);
                                        let imageList = images.current;
                                        if (imageList.length < 1) {
                                            throw Error('相片数量过少');
                                        }

                                        imageIndex.current -= 1;
                                        if (imageIndex.current <= 0) {
                                            imageIndex.current = images.current.length - 1;
                                        }

                                        let image = await getImageDebug(imageList[imageIndex.current].id);
                                        setImageMeta(image);
                                    } catch (error) {
                                        console.error(error);
                                    }
                                    setLoadingImage(false);
                                }}
                            >
                                Get Last Image
                                {loadingImage && <LoadingIcon size="1em" />}
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    // const axiosClient = axios.create({
                                    // 	withCredentials: true,
                                    // });

                                    const image = await (
                                        await fetch(imageMeta?.photoUrl || '', { credentials: 'omit', cache: 'no-cache' })
                                    ).blob();

                                    console.log('开始转换', image);

                                    // await heic2any({ blob: image, toType: 'image/png' }).then((imgFile) => {
                                    // 	let url = URL.createObjectURL(imgFile as Blob);
                                    // 	setImageMeta({
                                    // 		...imageMeta!,
                                    // 		thumbnailUrl: url,
                                    // 	});

                                    // 	console.log('转换渲染成功', imgFile);
                                    // });
                                }}
                            >
                                Get Current Image
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={() => {
                                    getFileOpenHandler();
                                }}
                            >
                                File Open Handler
                            </button>

                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    const currentDevice = await lzcAPI.current?.currentDevice
                                    console.log('currentDevice', currentDevice)

                                    let assets: any[] = [];
                                    try {
                                        const device = currentDevice!;
                                        let assetStats = device.photolibrary.ListAssetStats({
                                            isFilterVideo: false
                                        });

                                        await assetStats.forEach(async (item) => {
                                            assets.push(item);
                                        });

                                        console.log('assets', assets);
                                    } catch (error) {
                                        throw error;
                                    }
                                }}
                            >
                                getCurrentDevice
                            </button>
                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    // await AppCommon.Vibrate(VibrateType.EFFECT_DOUBLE_CLICK)
                                }}
                            >
                                振动测试
                            </button>
                            <button
                                style={{ marginTop: 10 }}
                                onClick={async () => {
                                    const currentDevice = await lzcAPI.current?.currentDevice
                                    console.log('currentDevice', currentDevice)

                                    let assets: any[] = [];
                                    try {
                                        const device = currentDevice!;
                                        const assetStats = device.photolibrary.ListAssetStats({
                                            isFilterVideo: false
                                        });

                                        await assetStats.forEach(async (item) => {
                                            assets.push(item);
                                        });
                                        console.log('assets', assets);
                                        let photosMeta: any[] = [];
                                        let ids = assets.map((item) => item.id)
                                        console.log('ids', ids);
                                        const assetsList = device.photolibrary.ListAssetsByIds({
                                            ids,
                                            sortBy: ListAssetsSortType.ASSETS_CREATE_DATE_ASC
                                        })

                                        await assetsList.forEach((item) => {
                                            photosMeta.push(item)
                                        })
                                        console.log('photosMeta', photosMeta);

                                    } catch (error) {
                                        throw error;
                                    }
                                }}
                            >
                                Get ListAssetStats
                            </button>
                        </div>

                        <p>
                            {/* 当前网络状态: <code>{netstate}</code> */}
                        </p>
                        {notImageUrl ? (
                            <p>
                                第一个相册图片为空数量: <code>{notImageUrl}</code>
                            </p>
                        ) : null}
                        {imageMeta != null && (
                            <AssetMetaReview meta={imageMeta} deletePhoto={deletePhoto} />
                        )}

                    </div>
                    <p className="read-the-docs">Click on the Vite and React logos to learn more</p>
                </>
            )}
        </div>
    );
}

export default App;
