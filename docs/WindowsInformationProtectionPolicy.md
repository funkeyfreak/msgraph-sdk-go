# WindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysWithoutContactBeforeUnenroll** | **int32** | Offline interval before app data is wiped (days)  | [optional] 
**MdmEnrollmentUrl** | **string** | Enrollment url for the MDM | [optional] 
**MinutesOfInactivityBeforeDeviceLock** | **int32** | Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 &lt;&#x3D; X &lt;&#x3D; 999. | [optional] 
**NumberOfPastPinsRemembered** | **int32** | Integer value that specifies the number of past PINs that can be associated to a user account that can&#39;t be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PasswordMaximumAttemptCount** | **int32** | The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 &lt;&#x3D; X &lt;&#x3D; 16 for desktop and 0 &lt;&#x3D; X &lt;&#x3D; 999 for mobile devices. | [optional] 
**PinExpirationDays** | **int32** | Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user&#39;s PIN will never expire. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PinLowercaseLetters** | [**interface{}**](.md) | Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinMinimumLength** | **int32** | Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest. | [optional] 
**PinSpecialCharacters** | [**interface{}**](.md) | Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! \&quot; # $ % &amp; &#39; ( ) * + , - . / : ; &lt; &#x3D; &gt; ? @ [ \\ ] ^ _ &#x60; { | } ~. Default is NotAllow. | [optional] 
**PinUppercaseLetters** | [**interface{}**](.md) | Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**RevokeOnMdmHandoffDisabled** | **bool** | New property in RS2, pending documentation | [optional] 
**WindowsHelloForBusinessBlocked** | **bool** | Boolean value that sets Windows Hello for Business as a method for signing into Windows. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


