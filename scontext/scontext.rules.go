// Code generated by protoc-gen-lzc. DO NOT EDIT.

package scontext

func ListPermission(methodName string) []int {
	switch methodName {
	case "cloud.lazycat.apis.sys.PackageManager.Install":
		// [INSTALL_PACKAGE]
		return []int{1}
	case "cloud.lazycat.apis.sys.PackageManager.Uninstall":
		// [INSTALL_PACKAGE]
		return []int{1}
	case "cloud.lazycat.apis.common.PeripheralDeviceService.MountFilesystem":
		// [USER_DOCUMENT]
		return []int{0}
	}
	return nil
}
