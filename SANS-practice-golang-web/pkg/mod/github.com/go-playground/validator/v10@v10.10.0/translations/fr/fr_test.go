package fr

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	french "github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	fre := french.New()
	uni := ut.New(fre, fre)
	trans, _ := uni.GetTranslator("fr")

	validate := validator.New()

	err := RegisterDefaultTranslations(validate, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string    `validate:"required"`
		RequiredNumber    int       `validate:"required"`
		RequiredMultiple  []string  `validate:"required"`
		LenString         string    `validate:"len=1"`
		LenNumber         float64   `validate:"len=1113.00"`
		LenMultiple       []string  `validate:"len=7"`
		MinString         string    `validate:"min=1"`
		MinNumber         float64   `validate:"min=1113.00"`
		MinMultiple       []string  `validate:"min=7"`
		MaxString         string    `validate:"max=3"`
		MaxNumber         float64   `validate:"max=1113.00"`
		MaxMultiple       []string  `validate:"max=7"`
		EqString          string    `validate:"eq=3"`
		EqNumber          float64   `validate:"eq=2.33"`
		EqMultiple        []string  `validate:"eq=7"`
		NeString          string    `validate:"ne="`
		NeNumber          float64   `validate:"ne=0.00"`
		NeMultiple        []string  `validate:"ne=0"`
		LtString          string    `validate:"lt=3"`
		LtNumber          float64   `validate:"lt=5.56"`
		LtMultiple        []string  `validate:"lt=2"`
		LtTime            time.Time `validate:"lt"`
		LteString         string    `validate:"lte=3"`
		LteNumber         float64   `validate:"lte=5.56"`
		LteMultiple       []string  `validate:"lte=2"`
		LteTime           time.Time `validate:"lte"`
		GtString          string    `validate:"gt=3"`
		GtNumber          float64   `validate:"gt=5.56"`
		GtMultiple        []string  `validate:"gt=2"`
		GtTime            time.Time `validate:"gt"`
		GteString         string    `validate:"gte=3"`
		GteNumber         float64   `validate:"gte=5.56"`
		GteMultiple       []string  `validate:"gte=2"`
		GteTime           time.Time `validate:"gte"`
		EqFieldString     string    `validate:"eqfield=MaxString"`
		EqCSFieldString   string    `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string    `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string    `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string    `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string    `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string    `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string    `validate:"nefield=EqFieldString"`
		GtFieldString     string    `validate:"gtfield=MaxString"`
		GteFieldString    string    `validate:"gtefield=MaxString"`
		LtFieldString     string    `validate:"ltfield=MaxString"`
		LteFieldString    string    `validate:"ltefield=MaxString"`
		AlphaString       string    `validate:"alpha"`
		AlphanumString    string    `validate:"alphanum"`
		NumericString     string    `validate:"numeric"`
		NumberString      string    `validate:"number"`
		HexadecimalString string    `validate:"hexadecimal"`
		HexColorString    string    `validate:"hexcolor"`
		RGBColorString    string    `validate:"rgb"`
		RGBAColorString   string    `validate:"rgba"`
		HSLColorString    string    `validate:"hsl"`
		HSLAColorString   string    `validate:"hsla"`
		Email             string    `validate:"email"`
		URL               string    `validate:"url"`
		URI               string    `validate:"uri"`
		Base64            string    `validate:"base64"`
		Contains          string    `validate:"contains=purpose"`
		ContainsAny       string    `validate:"containsany=!@#$"`
		Excludes          string    `validate:"excludes=text"`
		ExcludesAll       string    `validate:"excludesall=!@#$"`
		ExcludesRune      string    `validate:"excludesrune=???"`
		ISBN              string    `validate:"isbn"`
		ISBN10            string    `validate:"isbn10"`
		ISBN13            string    `validate:"isbn13"`
		UUID              string    `validate:"uuid"`
		UUID3             string    `validate:"uuid3"`
		UUID4             string    `validate:"uuid4"`
		UUID5             string    `validate:"uuid5"`
		ULID              string    `validate:"ulid"`
		ASCII             string    `validate:"ascii"`
		PrintableASCII    string    `validate:"printascii"`
		MultiByte         string    `validate:"multibyte"`
		DataURI           string    `validate:"datauri"`
		Latitude          string    `validate:"latitude"`
		Longitude         string    `validate:"longitude"`
		SSN               string    `validate:"ssn"`
		IP                string    `validate:"ip"`
		IPv4              string    `validate:"ipv4"`
		IPv6              string    `validate:"ipv6"`
		CIDR              string    `validate:"cidr"`
		CIDRv4            string    `validate:"cidrv4"`
		CIDRv6            string    `validate:"cidrv6"`
		TCPAddr           string    `validate:"tcp_addr"`
		TCPAddrv4         string    `validate:"tcp4_addr"`
		TCPAddrv6         string    `validate:"tcp6_addr"`
		UDPAddr           string    `validate:"udp_addr"`
		UDPAddrv4         string    `validate:"udp4_addr"`
		UDPAddrv6         string    `validate:"udp6_addr"`
		IPAddr            string    `validate:"ip_addr"`
		IPAddrv4          string    `validate:"ip4_addr"`
		IPAddrv6          string    `validate:"ip6_addr"`
		UinxAddr          string    `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string    `validate:"mac"`
		IsColor           string    `validate:"iscolor"`
		StrPtrMinLen      *string   `validate:"min=10"`
		StrPtrMaxLen      *string   `validate:"max=1"`
		StrPtrLen         *string   `validate:"len=2"`
		StrPtrLt          *string   `validate:"lt=1"`
		StrPtrLte         *string   `validate:"lte=1"`
		StrPtrGt          *string   `validate:"gt=10"`
		StrPtrGte         *string   `validate:"gte=10"`
		OneOfString       string    `validate:"oneof=red green"`
		OneOfInt          int       `validate:"oneof=5 63"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)

	test.LtString = "1234"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteTime = time.Now().Add(time.Hour * 24)

	test.LtFieldString = "12345"
	test.LteFieldString = "12345"

	test.LtCSFieldString = "1234"
	test.LteCSFieldString = "1234"

	test.AlphaString = "abc3"
	test.AlphanumString = "abc3!"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ???"

	test.ASCII = "????????????"
	test.PrintableASCII = "????????????"

	test.MultiByte = "1234feerf"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(validator.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.IsColor",
			expected: "IsColor doit ??tre une couleur valide",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC doit contenir une adresse MAC valide",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr doit ??tre une adresse IP r??solvable",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 doit ??tre une adresse IPv4 r??solvable",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 doit ??tre une adresse IPv6 r??solvable",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr doit ??tre une adressse UDP valide",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 doit ??tre une adressse IPv4 UDP valide",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 doit ??tre une adressse IPv6 UDP valide",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr doit ??tre une adressse TCP valide",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 doit ??tre une adressse IPv4 TCP valide",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 doit ??tre une adressse IPv6 TCP valide",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR doit contenir une notation CIDR valide",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 doit contenir une notation CIDR valide pour une adresse IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 doit contenir une notation CIDR valide pour une adresse IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN doit ??tre un num??ro SSN valide",
		},
		{
			ns:       "Test.IP",
			expected: "IP doit ??tre une adressse IP valide",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 doit ??tre une adressse IPv4 valide",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 doit ??tre une adressse IPv6 valide",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI doit contenir une URI data valide",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude doit contenir des coordonn??es latitude valides",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude doit contenir des coordonn??es longitudes valides",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte doit contenir des caract??res multioctets",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII ne doit contenir que des caract??res ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII ne doit contenir que des caract??res ascii affichables",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID doit ??tre un UUID valid",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 doit ??tre un UUID version 3 valid",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 doit ??tre un UUID version 4 valid",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 doit ??tre un UUID version 5 valid",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID doit ??tre une ULID valide",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN doit ??tre un num??ro ISBN valid",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 doit ??tre un num??ro ISBN-10 valid",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 doit ??tre un num??ro ISBN-13 valid",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes ne doit pas contenir le texte 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll ne doit pas contenir l'un des caract??res suivants '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune ne doit pas contenir ce qui suit '???'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny doit contenir au moins l' un des caract??res suivants '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains doit contenir le texte 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 doit ??tre une cha??ne de caract??res au format Base64 valide",
		},
		{
			ns:       "Test.Email",
			expected: "Email doit ??tre une adresse email valide",
		},
		{
			ns:       "Test.URL",
			expected: "URL doit ??tre une URL valide",
		},
		{
			ns:       "Test.URI",
			expected: "URI doit ??tre une URI valide",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString doit ??tre une couleur au format RGB valide",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString doit ??tre une couleur au format RGBA valide",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString doit ??tre une couleur au format HSL valide",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString doit ??tre une couleur au format HSLA valide",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString doit ??tre une cha??ne de caract??res au format hexad??cimal valide",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString doit ??tre une couleur au format HEX valide",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString doit ??tre un nombre valid",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString doit ??tre une valeur num??rique valide",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString ne doit contenir que des caract??res alphanum??riques",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString ne doit contenir que des caract??res alphab??tiques",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString doit ??tre inf??rieur ?? MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString doit ??tre inf??rieur ou ??gal ?? MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString doit ??tre sup??rieur ?? MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString doit ??tre sup??rieur ou ??gal ?? MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString ne doit pas ??tre ??gal ?? EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString doit ??tre inf??rieur ?? Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString doit ??tre inf??rieur ou ??gal ?? Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString doit ??tre sup??rieur ?? Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString doit ??tre sup??rieur ou ??gal ?? Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString ne doit pas ??tre ??gal ?? Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString doit ??tre ??gal ?? Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString doit ??tre ??gal ?? MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString doit faire une taille d'au moins 3 caract??res",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber doit ??tre 5,56 ou plus",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple doit contenir au moins 2 elements",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime doit ??tre apr??s ou pendant la date et l'heure actuelle",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString doit avoir une taille sup??rieur ?? 3 caract??res",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber doit ??tre sup??rieur ?? 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple doit contenir plus de 2 elements",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime doit ??tre apr??s la date et l'heure actuelle",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString doit faire une taille maximum de 3 caract??res",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber doit faire 5,56 ou moins",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple doit contenir un maximum de 2 elements",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime doit ??tre avant ou pendant la date et l'heure actuelle",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString doit avoir une taille inf??rieure ?? 3 caract??res",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber doit ??tre inf??rieur ?? 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple doit contenir mois de 2 elements",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime doit ??tre avant la date et l'heure actuelle",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString ne doit pas ??tre ??gal ?? ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber ne doit pas ??tre ??gal ?? 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple ne doit pas ??tre ??gal ?? 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString n'est pas ??gal ?? 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber n'est pas ??gal ?? 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple n'est pas ??gal ?? 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString doit faire une taille maximum de 3 caract??res",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber doit ??tre ??gal ?? 1???113,00 ou moins",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple doit contenir au maximum 7 elements",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString doit faire une taille minimum de 1 caract??re",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber doit ??tre ??gal ?? 1???113,00 ou plus",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple doit contenir au moins 7 elements",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString doit faire une taille de 1 caract??re",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber doit ??tre ??gal ?? 1???113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple doit contenir 7 elements",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString est un champ obligatoire",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber est un champ obligatoire",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple est un champ obligatoire",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen doit faire une taille minimum de 10 caract??res",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen doit faire une taille maximum de 1 caract??re",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen doit faire une taille de 2 caract??res",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt doit avoir une taille inf??rieure ?? 1 caract??re",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte doit faire une taille maximum de 1 caract??re",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt doit avoir une taille sup??rieur ?? 10 caract??res",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte doit faire une taille d'au moins 10 caract??res",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString doit ??tre l'un des choix suivants [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt doit ??tre l'un des choix suivants [5 63]",
		},
	}

	for _, tt := range tests {

		var fe validator.FieldError

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}

}
