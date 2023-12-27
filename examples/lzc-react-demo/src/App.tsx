import React, { useState, useEffect, useRef } from 'react';
import reactLogo from './assets/react.svg';
import './App.css';
import { LoadingIcon } from 'tdesign-icons-react';
// import { firstValueFrom } from 'rxjs';
import { lzcAPIGateway, EndDeviceProxy } from '@lazycatcloud/sdk/index';
import { ListAlbumsReply, PhotoMeta } from '@lazycatcloud/sdk/localdevice/photo';
import { DeviceInfo } from '@lazycatcloud/sdk/localdevice/device';
import { Permission } from '@lazycatcloud/sdk/localdevice/permission';

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
            device.photolibrary.DeletePhoto({
                id: ids,
            });
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

    return (
        <div className="App">
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
                    <button
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
                    <div style={{ width: 20 }}></div>
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
                        Get All Album
                    </button>

                    <button
                        style={{ marginTop: 10 }}
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

                                if (imageIndex.current > imageList.length - 1) {
                                    imageIndex.current = 0;
                                }
                                let image = await getImageDebug(imageList[imageIndex.current].id);
                                // let image = await getImageDebug('1FC00A27-FF03-48DF-9DBB-5F30DCB26FAD/L0/001'); // IMG_0009.MOV
                                // let image = await getImageDebug('32354'); // IMG_0088.HEIC
                                setImageMeta(image);
                                imageIndex.current += 1;

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

                    {/* <button
						style={{ marginTop: 10 }}
						onClick={() => {
							putPhoto();
						}}
					>
						Save Current Image
					</button> */}

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
                            console.log(await lzcAPI.current?.currentDevice)
                        }}
                    >
                        getCurrentDevice
                    </button>
                </div>

                <p>
                    当前网络状态: <code>{netState}</code>
                </p>
                {notImageUrl ? (
                    <p>
                        第一个相册图片为空数量: <code>{notImageUrl}</code>
                    </p>
                ) : null}

                {imageMeta != null && (
                    <div>
                        查询图片 <span style={{ width: 100 }}>({imageMeta.id})</span>:
                        <a href={imageMeta.photoUrl} target={'_blank'}>
                            查看原图
                        </a>
                        {'  '}
                        <a href={imageMeta.thumbnailUrl} target={'_blank'}>
                            缩略图
                        </a>
                        {'  '}
                        <a
                            onClick={() => {
                                deletePhoto([imageMeta.id]);
                            }}
                        >
                            删除图片
                        </a>{' '}
                        <p>{imageMeta.fileName}</p>
                        <div style={{ padding: 2, background: '#0001' }}>
                            <img style={{ width: 150 }} src={imageMeta.thumbnailUrl || imageMeta.photoUrl} />
                        </div>
                    </div>
                )}
            </div>
            <p className="read-the-docs">Click on the Vite and React logos to learn more</p>
        </div>
    );
}

export default App;
