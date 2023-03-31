//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package jbig2 ;import (_b "github.com/unidoc/unipdf/v3/internal/bitwise";_ge "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_gg "github.com/unidoc/unipdf/v3/internal/jbig2/document";_c "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";
_ba "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_g "sort";);func (_gee Globals )ToDocumentGlobals ()*_gg .Globals {if _gee ==nil {return nil ;};_ea :=[]*_c .Header {};for _ ,_cee :=range _gee {_ea =append (_ea ,_cee );};_g .Slice (_ea ,func (_gc ,_bef int )bool {return _ea [_gc ].SegmentNumber < _ea [_bef ].SegmentNumber });
return &_gg .Globals {Segments :_ea };};type Globals map[int ]*_c .Header ;func DecodeBytes (encoded []byte ,parameters _ge .Parameters ,globals ...Globals )([]byte ,error ){var _f Globals ;if len (globals )> 0{_f =globals [0];};_bd ,_eb :=_ge .Decode (encoded ,parameters ,_f .ToDocumentGlobals ());
if _eb !=nil {return nil ,_eb ;};return _bd .DecodeNextPage ();};func DecodeGlobals (encoded []byte )(Globals ,error ){const _gge ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";_a :=_b .NewReader (encoded );_d ,_ce :=_gg .DecodeDocument (_a ,nil );
if _ce !=nil {return nil ,_ba .Wrap (_ce ,_gge ,"");};if _d .GlobalSegments ==nil ||(_d .GlobalSegments .Segments ==nil ){return nil ,_ba .Error (_gge ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");
};_df :=Globals {};for _ ,_be :=range _d .GlobalSegments .Segments {_df [int (_be .SegmentNumber )]=_be ;};return _df ,nil ;};