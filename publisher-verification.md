# Process of Publisher Verification

Publications are signed by publishers. To trust a publication signature, the publisher has to be trusted. This document describes how trust is established in various scenarios.


## Scenario 1: No Domain Security Service And No CA

The publication address is secured by the message bus using the publisher's account. This is the simplest scenario and acceptable for local usage. 

In this case, consumers always accept the publication by the publisher as only the publisher can publish on its address.


## Scenario 2: No Domain Security Service; Using A CA

Publishers carry a certificate signed by a CA (Lets Encrypt)

In this case, the publication of the publisher identity contains a certificate from Lets Encrypt. The publication is accepted if the certificate is verified by Lets Encrypt CA public key. Publishers require a copy of the Lets Encrypt public key. 

This can be downloaded manually or automatically on startup from the Lets Encrypt website using SSL with domain verification. The latter requires a trusted DNS and is still vulnerable to DNS cache poisening.

## Scenario 3: With Domain Security Service; Not Yet Joined 

A secure domain configuration is set in the messenger configuration so the publisher knows to require DSS signed publisher discovery.

Before the publisher joins the secure domain, it publishes its identity without a DSS signature. Subscribers notice the lack of DSS signature and mark the publisher as untrusted (or discard it entirely).

The DSS will store the publisher identity for review and approval by the administrator. When approved, the DSS will issue an updated identity with signature to the publisher. The publisher will then have joined the secure domain. (next scenario).

## Scenario 4: Receiving the Domain Security Service Identity 

In DSS secured domains, consumers first subscribe to the DSS identity. However, who verifies the identity of the DSS itself? This comes down to scenario 1 (address protection) or 2 (certificate using a CA).

Messenger configuration for these scenarios are:
{
  dssVerification: address | letsencrypt.com
}

'address' means that publication address is secured
'letsencrypt.com' the DSS contains a certificate which is verified against Lets Encrypt

In the latter case if no certificate is provided or the verification failed, the publisher is marked as untrusted.


## Scenario 5: With Domain Security Service; Joined

In secured domains, consumers first subscribe to the DSS. Once the DSS identity is received, the user subscribes to publishers. This ensures that the DSS identity exists when a publisher is discovered.

The published identity contains a DSS signature. If the signature verifies with the DSS public key and the JSON encoded identity as payload, then the publisher is marked as verified.


## Scenario 6: Updating the DSS signature

The DSS keys are regularly refreshed. The DSS re-issues identity signatures for all publishers after this refresh.

