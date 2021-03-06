package tags_test

import (
	"fmt"

	"github.com/xgfone/go-tools/tags"
)

func ExampleTag() {
	type TagTest struct {
		F1 string `tag1:"123" tag2:"456" tag3:"789" tag4:"000"`
		F2 string `tag1:"aaa" tag2:"bbb" tag3:"ccc" tag5:"zzz"`
		F3 string `tag1:"ddd" tag2:"eee" tag3:"fff" tag6:"yyy"`
	}
	// tags.Debug = true
	tag := tags.NewTag(TagTest{})
	// tag.Build() // Suggest to use this method firstly.
	tag.BuildTags([]string{"tag1", "tag2"}).BuildTag("tag5").BuildTag("tag6")
	// fmt.Println(tag.Audit())

	// // Get
	fmt.Println("--------- Get")
	fmt.Println(tag.Get("tag1"))
	fmt.Println(tag.Get("tag2"))
	fmt.Println(tag.Get("tag4"))

	// // GetToField
	fmt.Println("--------- GetToField")
	for _, t := range tag.GetToField("tag1") {
		fmt.Println(t.Field, t.Value)
	}
	fmt.Println(tag.GetToField("tag7"))

	// GetAllByField
	fmt.Println("--------- GetAllByField")
	for _, tv := range tag.GetAllByField("F1") {
		fmt.Println(tv.Tag, tv.Value)
	}
	fmt.Println(tag.GetAllByField("F4"))

	// // GetWithField
	fmt.Println("--------- GetWithField")
	fmt.Println(tag.GetWithField("tag1"))
	field, value := tag.GetWithField("tag4")
	fmt.Println(field, value, "End") // End is for saving the output whitespaces.
	fmt.Println(tag.GetWithField("tag5"))
	fmt.Println(tag.GetWithField("tag6"))

	// // GetByField
	fmt.Println("--------- GetByField")
	fmt.Println(tag.GetByField("tag1", "F1"))
	fmt.Println(tag.GetByField("tag1", "F2"))
	fmt.Println(tag.GetByField("tag4", "F2"))

	// GetAll
	fmt.Println("--------- GetAll")
	for _, ft := range tag.GetAll() {
		fmt.Println(ft.Field, ft.Tag, ft.Value)
	}

	// TravelByTag
	fmt.Println("--------- TravelByTag")
	tag.TravelByTag("tag1", func(field, value string) {
		fmt.Println(field, value)
	})

	// TravelByField
	fmt.Println("--------- TravelByField")
	tag.TravelByField("F1", func(tag, value string) {
		fmt.Println(tag, value)
	})

	// Output:
	// --------- Get
	// 123
	// 456
	//
	// --------- GetToField
	// F1 123
	// F2 aaa
	// F3 ddd
	// []
	// --------- GetAllByField
	// tag1 123
	// tag2 456
	// []
	// --------- GetWithField
	// F1 123
	//   End
	// F2 zzz
	// F3 yyy
	// --------- GetByField
	// 123
	// aaa
	//
	// --------- GetAll
	// F1 tag1 123
	// F1 tag2 456
	// F2 tag1 aaa
	// F2 tag2 bbb
	// F2 tag5 zzz
	// F3 tag1 ddd
	// F3 tag2 eee
	// F3 tag6 yyy
	// --------- TravelByTag
	// F1 123
	// F2 aaa
	// F3 ddd
	// --------- TravelByField
	// tag1 123
	// tag2 456
}
