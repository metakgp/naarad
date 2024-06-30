# Naarad Subscription Instruction



## User Registration

1. Visit [Naarad Signup](https://naarad.metakgp.org/signup) page
2. You will be prompted to enter your institute email
3. Once email is entered, you will receive a verification OTP on the provided email
4. Enter the OTP in the available field and click the `Verify` button
5. After a while, you will then receive a prompt on the screen saying `Credentials sent to your institute email`
6. Check your institute email for credentials and make sure to keep them safe for future use
7. Follow through [webapp user login](#user-login) step

| ![image](https://github.com/metakgp/naarad/assets/86282911/f4e828c6-fa58-496b-bf09-7cb6cd7d2b09) | ![image](https://github.com/metakgp/naarad/assets/86282911/94e0acad-d05c-4640-a75d-debb3276abcc) |
| ---------------------------------- | -------------------------------- |
| ![image](https://github.com/metakgp/naarad/assets/86282911/1aaa2595-15c1-4530-8c91-e06c1790382f) | ![image](https://github.com/metakgp/naarad/assets/86282911/5af29554-8b3a-4776-81fb-26a9af41128b) |

## Webapp

### User Login

0. After following all the steps for [registering a user on naarad](#user-registration) you will be redirected to [Naarad Login](https://naarad.metakgp.org/login) page
1. You will be prompted to enter username and password, enter those which you just received on your institute email
2. Finally, press the `login` button, now you are authorised to access naarad and subscribe to limited topics

<div align="center">
  <img src="https://github.com/metakgp/naarad/assets/86282911/ab48f663-4f2c-41a5-95e6-c02a845cb368">
</div>

### Configuring Notification

Make sure to grant the webapp notification permission.

<div align="center">
  <img src="https://github.com/metakgp/naarad/assets/86282911/1b0115cb-0d79-4f70-99c5-7adeb2e92a9c">
</div>

#### Enabling WebPush

In order to receive notifications even if the [webapp's tab](https://naarad.metakgp.org) is closed, follow the steps below:

1. Click on `Settings` button on the left panel
2. Enable `Background Notifications` (4th option from top)

> [!Tip]
> It is recommeded to also change the `Delete Notifications` setting to `Never` in order to store old notifications

<div align="center">
  <img src="https://github.com/metakgp/naarad/assets/86282911/45ab7ef1-91d4-4fea-8375-6b0c70c4b9c4">
</div>

### Subscribing to MFTP

#### Automatic

After completing [Naarad User Login](#user-login), visit https://naarad.metakgp.org/kgp-mftp and you will be automatically subscribed to MFTP aka CDC Noticeboard.

#### Manual

If the [automatic](#automatic) method doesn't work, then you can do it manually as described below:

1. Click on `+ Subscribe to topic` button on the left panel
2. Enter `kgp-mftp` in the text box and click `Subscrbe` button
3. In case you were not logged in, it will require for credentials, enter the ones you received during [user registration](#user-registration)

<div align="center">
  <img src="https://github.com/metakgp/naarad/assets/86282911/43d66861-b5ef-4776-a940-21565da7b5a0">
</div>

## Mobile Application (Android, iOS)

### Download NTFY application

<div class="mt-8 flex flex-wrap gap-x-3 gap-y-4">
  <a target="_blank" href="https://play.google.com/store/apps/details?id=io.heckel.ntfy">
    <img alt="Get it on Google Play" src="https://ntfy.sh/_next/static/media/badge-google.19268080.png" width="168" height="50" decoding="async" data-nimg="1" loading="lazy" style="color:transparent">
  </a>
  <a target="_blank" href="https://f-droid.org/en/packages/io.heckel.ntfy/">
    <img alt="Get it on F-Droid" src="https://ntfy.sh/_next/static/media/badge-fdroid.f6ae6646.png" width="168" height="50" decoding="async" data-nimg="1" loading="lazy" style="color:transparent">
  </a>
  <a target="_blank" href="https://apps.apple.com/us/app/ntfy/id1625396347">
    <img alt="Download on the App Store" src="https://ntfy.sh/_next/static/media/badge-apple.4bec723d.png" width="148" height="50" decoding="async" data-nimg="1" loading="lazy" style="color:transparent">
  </a>
</div>

### Subscribing to MFTP

1. Click on `+` icon in bottom right corner
2. Enter topic name as: `kgp-mftp`
3. Check `Use another server`
4. Replace `https://ntfy.sh` with `https://naarad.metakgp.org`
5. Press the `Subscribe` button
6. It will now prompt you to Login, enter the credentials received during [user registration](#user-registration)
7. Click the `Login` button

| ![](https://github.com/metakgp/naarad/assets/86282911/2e6cb3df-d65b-41a0-bcf0-7623607bf56b) | ![](https://github.com/metakgp/naarad/assets/86282911/0a65338a-705d-47dd-8fa7-e399e23f4908) |
| ---------------------------------- | -------------------------------- |
