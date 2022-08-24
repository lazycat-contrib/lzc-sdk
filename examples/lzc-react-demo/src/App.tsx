import React, { useState, useEffect, useRef } from 'react';
import reactLogo from './assets/react.svg';
import './App.css';
import { LoadingIcon } from 'tdesign-icons-react';
import { firstValueFrom } from 'rxjs';

import { lzcAPIGateway, EndDeviceProxy } from '@lazycatcloud/sdk';
import { ListAlbumsReply, PhotoMeta } from '@lazycatcloud/sdk/dist/localdevice/photo';

function App() {
	const [loading, setLoading] = useState(false);
	const [netState, setNetState] = useState('无网络');
	const lzcAPI = useRef<lzcAPIGateway>();
	const currentDevice = useRef<EndDeviceProxy>();

	const [imageMeta, setImageMeta] = useState<PhotoMeta>();
	const images = useRef<PhotoMeta[]>(Array<PhotoMeta>());
	const imageIndex = useRef(0);
	const [loadingImage, setLoadingImage] = useState(false);

	useEffect(() => {
		if (!lzcAPI.current) {
			lzcAPI.current = new lzcAPIGateway();

			// 调试
			getAlbumsDebug();
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
			currentDevice.current = (await lzcAPI.current?.currentDevice)!;
		}
		return currentDevice.current;
	};

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
									images.current = await getAlbumDebug([]);
								}
								let imageList = images.current;
								if (imageList.length < 1) {
									throw Error('相片数量过少');
								}

								if (imageIndex.current > imageList.length - 1) {
									imageIndex.current = 0;
								}
								let image = await getImageDebug(imageList[imageIndex.current].id);
								setImageMeta(imageList[imageIndex.current]);
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
				</div>

				<p>
					当前网络状态: <code>{netState}</code>
				</p>

				{imageMeta != null && (
					<div>
						查询图片 <span style={{ width: 100 }}>({imageMeta.id})</span>:
						<a href={imageMeta.photoUrl} target={'_blank'}>
							查看原图
						</a>
						{'  '}
						<a
							onClick={() => {
								deletePhoto([imageMeta.id]);
							}}
						>
							删除图片
						</a>{' '}
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
