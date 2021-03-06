// kibosh
//
// Copyright (c) 2017-Present Pivotal Software, Inc. All Rights Reserved.
//
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License”); you may not use this file except in compliance with the License. You may
// obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing permissions and
// limitations under the License.

package broker_test

import (
	b64 "encoding/base64"

	. "github.com/cf-platform-eng/kibosh/pkg/broker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cluster Params", func() {
	Context("extract parameters", func() {
		It("properly extracts config parameters", func() {
			json := []byte(`{"clusterConfig" : {"server": "server url", "token":"token data", "certificateAuthorityData":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURKRENDQWd5Z0F3SUJBZ0lVVHltSk1uWEU0aXp5QlBvalM1enpXU2tzNmVrd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0RURUxNQWtHQTFVRUF4TUNZMkV3SGhjTk1UZ3dOVEE0TVRneU16SXdXaGNOTVRrd05UQTRNVGd5TXpJdwpXakFOTVFzd0NRWURWUVFERXdKallUQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCCkFLNGd3ZnFpeG1KT1JjSmtERjZVdmNVQWl5TEdlRks1Z0JnaEU3MVFzMnZ4UU4vT1AxekVHMkVSWHZheFIzbUYKVFdxTkNlTTl5d1Fpcm9FTmtFODd2LzFBZXZudkJQMDVZczdmaU5pS0ZNZTRYV091UWRlNXR0S3JpdlpJRWtCawpTT2psdXlQR0g4d3JTY0J1alZQelQvMGxLR3FKUW1iTGFTVm1qczczK0NINFpEZ2ZILzN0c0tpQTRaWU54Z2JGCnY0WWRnTFJHSkdTZjN5NlhyaWoxaVpMaUdhYjVWbDVLUSs2T0ZqR0wxbEEybyt4SGV2d2J0T2hQdlB1emNYbHkKd1RJay8rSEw0ckRMOG9Oemg5QTFldlFCaDJHRzhUSmFQMXVldVVXSXlHZENyb2FqTUZMN1ZaZkd1aUZLK2UyUwplODNYNHdzakJSc0t6RFlKL21IcjE5a0NBd0VBQWFOOE1Ib3dIUVlEVlIwT0JCWUVGSDdsRlJWSFZSeXFYQkhJClQ1cGdJVmRtM0JmcU1FZ0dBMVVkSXdSQk1EK0FGSDdsRlJWSFZSeXFYQkhJVDVwZ0lWZG0zQmZxb1JHa0R6QU4KTVFzd0NRWURWUVFERXdKallZSVVUeW1KTW5YRTRpenlCUG9qUzV6eldTa3M2ZWt3RHdZRFZSMFRBUUgvQkFVdwpBd0VCL3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQUs2SUtJdGJOMVFaR2pMWUZsTU1KbDJrcVFCNG9KekgyClZLczhxTCtMZDJHVHNIWDVxKzFhV2ZLcVRha0V1QVdwTGZYWUZEUXc3TXYyak9rZGQ0WEV6MXEwZ3k1QWw1MVYKZnlYaXBJalBMdFYwK21DdGVkc2hFNHJZVjZvQUp4RFE2MzJ2b3JpWVJpR3A3SHVqL254VjFMbmUwQzlQdmI0UAp1NVYrZGxQRHZSR3J2Y1dtNjk4bC9PQncyNk9GcHFCQytBUExteW5SMDBXL2xQQURHOWpaT0ZiblNlRGFPMkhqClcwQzhzb3QrYkZianFsaU01T2hBU0RwOFI2VHBqU1hWNEFqZzE5blMxM1M0bVZVSGFtOXJOTkw4aWVhdVdVMUUKdUVaUFBNb0hHcGlZQ29CelEwYmdqL0xaVDR1YzVlZ1Mrb29XdUJTKzM0Mk1KcVFFa2NVRWFBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}}`)

			cluster, exist := ExtractClusterConfig(json)

			Expect(exist).To(BeTrue())
			Expect(cluster.Server).To(Equal("server url"))
			Expect(cluster.Token).To(Equal("token data"))

			caData, _ := b64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURKRENDQWd5Z0F3SUJBZ0lVVHltSk1uWEU0aXp5QlBvalM1enpXU2tzNmVrd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0RURUxNQWtHQTFVRUF4TUNZMkV3SGhjTk1UZ3dOVEE0TVRneU16SXdXaGNOTVRrd05UQTRNVGd5TXpJdwpXakFOTVFzd0NRWURWUVFERXdKallUQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCCkFLNGd3ZnFpeG1KT1JjSmtERjZVdmNVQWl5TEdlRks1Z0JnaEU3MVFzMnZ4UU4vT1AxekVHMkVSWHZheFIzbUYKVFdxTkNlTTl5d1Fpcm9FTmtFODd2LzFBZXZudkJQMDVZczdmaU5pS0ZNZTRYV091UWRlNXR0S3JpdlpJRWtCawpTT2psdXlQR0g4d3JTY0J1alZQelQvMGxLR3FKUW1iTGFTVm1qczczK0NINFpEZ2ZILzN0c0tpQTRaWU54Z2JGCnY0WWRnTFJHSkdTZjN5NlhyaWoxaVpMaUdhYjVWbDVLUSs2T0ZqR0wxbEEybyt4SGV2d2J0T2hQdlB1emNYbHkKd1RJay8rSEw0ckRMOG9Oemg5QTFldlFCaDJHRzhUSmFQMXVldVVXSXlHZENyb2FqTUZMN1ZaZkd1aUZLK2UyUwplODNYNHdzakJSc0t6RFlKL21IcjE5a0NBd0VBQWFOOE1Ib3dIUVlEVlIwT0JCWUVGSDdsRlJWSFZSeXFYQkhJClQ1cGdJVmRtM0JmcU1FZ0dBMVVkSXdSQk1EK0FGSDdsRlJWSFZSeXFYQkhJVDVwZ0lWZG0zQmZxb1JHa0R6QU4KTVFzd0NRWURWUVFERXdKallZSVVUeW1KTW5YRTRpenlCUG9qUzV6eldTa3M2ZWt3RHdZRFZSMFRBUUgvQkFVdwpBd0VCL3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQUs2SUtJdGJOMVFaR2pMWUZsTU1KbDJrcVFCNG9KekgyClZLczhxTCtMZDJHVHNIWDVxKzFhV2ZLcVRha0V1QVdwTGZYWUZEUXc3TXYyak9rZGQ0WEV6MXEwZ3k1QWw1MVYKZnlYaXBJalBMdFYwK21DdGVkc2hFNHJZVjZvQUp4RFE2MzJ2b3JpWVJpR3A3SHVqL254VjFMbmUwQzlQdmI0UAp1NVYrZGxQRHZSR3J2Y1dtNjk4bC9PQncyNk9GcHFCQytBUExteW5SMDBXL2xQQURHOWpaT0ZiblNlRGFPMkhqClcwQzhzb3QrYkZianFsaU01T2hBU0RwOFI2VHBqU1hWNEFqZzE5blMxM1M0bVZVSGFtOXJOTkw4aWVhdVdVMUUKdUVaUFBNb0hHcGlZQ29CelEwYmdqL0xaVDR1YzVlZ1Mrb29XdUJTKzM0Mk1KcVFFa2NVRWFBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

			Expect(cluster.CADataRaw).To(Equal("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURKRENDQWd5Z0F3SUJBZ0lVVHltSk1uWEU0aXp5QlBvalM1enpXU2tzNmVrd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0RURUxNQWtHQTFVRUF4TUNZMkV3SGhjTk1UZ3dOVEE0TVRneU16SXdXaGNOTVRrd05UQTRNVGd5TXpJdwpXakFOTVFzd0NRWURWUVFERXdKallUQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCCkFLNGd3ZnFpeG1KT1JjSmtERjZVdmNVQWl5TEdlRks1Z0JnaEU3MVFzMnZ4UU4vT1AxekVHMkVSWHZheFIzbUYKVFdxTkNlTTl5d1Fpcm9FTmtFODd2LzFBZXZudkJQMDVZczdmaU5pS0ZNZTRYV091UWRlNXR0S3JpdlpJRWtCawpTT2psdXlQR0g4d3JTY0J1alZQelQvMGxLR3FKUW1iTGFTVm1qczczK0NINFpEZ2ZILzN0c0tpQTRaWU54Z2JGCnY0WWRnTFJHSkdTZjN5NlhyaWoxaVpMaUdhYjVWbDVLUSs2T0ZqR0wxbEEybyt4SGV2d2J0T2hQdlB1emNYbHkKd1RJay8rSEw0ckRMOG9Oemg5QTFldlFCaDJHRzhUSmFQMXVldVVXSXlHZENyb2FqTUZMN1ZaZkd1aUZLK2UyUwplODNYNHdzakJSc0t6RFlKL21IcjE5a0NBd0VBQWFOOE1Ib3dIUVlEVlIwT0JCWUVGSDdsRlJWSFZSeXFYQkhJClQ1cGdJVmRtM0JmcU1FZ0dBMVVkSXdSQk1EK0FGSDdsRlJWSFZSeXFYQkhJVDVwZ0lWZG0zQmZxb1JHa0R6QU4KTVFzd0NRWURWUVFERXdKallZSVVUeW1KTW5YRTRpenlCUG9qUzV6eldTa3M2ZWt3RHdZRFZSMFRBUUgvQkFVdwpBd0VCL3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQUs2SUtJdGJOMVFaR2pMWUZsTU1KbDJrcVFCNG9KekgyClZLczhxTCtMZDJHVHNIWDVxKzFhV2ZLcVRha0V1QVdwTGZYWUZEUXc3TXYyak9rZGQ0WEV6MXEwZ3k1QWw1MVYKZnlYaXBJalBMdFYwK21DdGVkc2hFNHJZVjZvQUp4RFE2MzJ2b3JpWVJpR3A3SHVqL254VjFMbmUwQzlQdmI0UAp1NVYrZGxQRHZSR3J2Y1dtNjk4bC9PQncyNk9GcHFCQytBUExteW5SMDBXL2xQQURHOWpaT0ZiblNlRGFPMkhqClcwQzhzb3QrYkZianFsaU01T2hBU0RwOFI2VHBqU1hWNEFqZzE5blMxM1M0bVZVSGFtOXJOTkw4aWVhdVdVMUUKdUVaUFBNb0hHcGlZQ29CelEwYmdqL0xaVDR1YzVlZ1Mrb29XdUJTKzM0Mk1KcVFFa2NVRWFBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="))
			Expect(cluster.CAData).To(Equal(caData))
		})

		It("handles empty config parameters", func() {
			json := []byte("")

			_, exist := ExtractClusterConfig(json)

			Expect(exist).To(BeFalse())
		})

		It("handles non cluster config", func() {
			json := []byte(`{"some":"config", "other":{"config":"here"}}`)

			_, exist := ExtractClusterConfig(json)

			Expect(exist).To(BeFalse())
		})
	})
})
