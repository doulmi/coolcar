package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwFE7+CbqDNyGEUiT6rph
nwFLJwV04H0FYVc/S+1tzK/PrwmBV3JuFfakU0DMsBVtuqkYhDSjASHuohvhgnq7
aDoUt1xlzfY+SRTzU1Vs+X4fiTCB66sAcPzxb0e5RBKyvBm45F44olqNiOrudSsM
BQ2L8XmrnefzNnkZdBr6kDL8ATR7/dqf23mNuTgI3G1zHfSNxlk2PtDLSVivZ8ZZ
TZsShAmQfvATajpKmZCmlfKfFBBLieEz2amzGR7CPOwofbCqzMbI7M1kgRuS4vBD 
q4FTe3OMzBu5PmjbuIOG2mQ0DhC5lBD5xnEF/k3QyXvuB9SFTwNKhauKheCPHTxK
uQIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

	if err != nil {
		t.Fatalf("Cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: publicKey,
	}

	// token will be expired at 1679576933
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNjc5NTc2OTMzfQ.ZXjGAjhOPRYJ9mdkUubFn2jrRvHNUVcrvvrcweNLqfB0ipstsEzDkoqKvtTdu5ydiXz07ns-u9IMchMktCZehUoAKENlo4vuiHd4dFsWM97sAXFQcdUeF36jdkmFreRjYH6fWFju979ukhAck5DJRPU_ZrDNY6jE88pWU2B1eu2Yx6XXINF_IXutYUhvJ3a4unYNpMDSHzaEFdozbxMElbLRq_xvyie0gjqUGUDm6LGcepUaDvemKOvkxa-UlyLeH3wcQGT_PzPImqh20LensIPLzHupDyxryPv_vTX-G7qPQ_VvleNte5b-lWgQJ8ObkpF71JV7_MQOOZBtgsqi8w"

	fakeToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNjc5NTc2OTMzdd.ZXjGAjhOPRYJ9mdkUubFn2jrRvHNUVcrvvrcweNLqfB0ipstsEzDkoqKvtTdu5ydiXz07ns-u9IMchMktCZehUoAKENlo4vuiHd4dFsWM97sAXFQcdUeF36jdkmFreRjYH6fWFju979ukhAck5DJRPU_ZrDNY6jE88pWU2B1eu2Yx6XXINF_IXutYUhvJ3a4unYNpMDSHzaEFdozbxMElbLRq_xvyie0gjqUGUDm6LGcepUaDvemKOvkxa-UlyLeH3wcQGT_PzPImqh20LensIPLzHupDyxryPv_vTX-G7qPQ_VvleNte5b-lWgQJ8ObkpF71JV7_MQOOZBtgsqi8w"

	cases := []struct {
		name        string
		token       string
		now         time.Time
		expected    string
		expectedErr bool
	}{
		{
			name:        "valid_token",
			token:       token,
			now:         time.Unix(1679570933, 0),
			expected:    "2",
			expectedErr: false,
		},
		{
			name:        "token_expired",
			token:       token,
			now:         time.Unix(1679579933, 0),
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "bad_token",
			token:       "bad_token",
			now:         time.Unix(1679576933, 0),
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "fake_token",
			token:       fakeToken,
			now:         time.Unix(1679576933, 0),
			expected:    "",
			expectedErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// TODO: this can't work, can't find any example
			jwt.WithTimeFunc(func() time.Time {
				fmt.Printf("is called: %v\n", c.now)
				return c.now
			})

			accountId, err := v.Verify(token)

			if c.expectedErr && err != nil {
				t.Errorf("Verification failed: %v", err)
			}

			if !c.expectedErr && err != nil {
				t.Errorf("expected an error, but got no error: %v", err)
			}

			assert.Equal(t, c.expected, accountId)
		})
	}
}
