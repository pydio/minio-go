package minio

import (
	"strings"
	"testing"
)

func TestClient_GetBucketTagging(t *testing.T) {
	response := `<Tagging>
           <TagSet>
              <Tag>
                <Key>Project</Key>
               <Value>Project One</Value>
              </Tag>
              <Tag>
                <Key>User</Key>
                <Value>jsmith</Value>
              </Tag>
           </TagSet>
         </Tagging>`
	reader := strings.NewReader(response)
	tagging := Tagging{}
	er := xmlDecoder(reader, &tagging)
	if er != nil {
		t.Fatalf("cannot decode response: %s", er.Error())
	}
	if len(tagging.TagSet.Tag) != 2 {
		t.Fatalf("wrong tags length")
	}
	tags := tagging.TagSet.Tag
	if tags[0].Key != "Project" || tags[0].Value != "Project One" {
		t.Fatalf("wrong tag[0] value %v", tags[0])
	}
}