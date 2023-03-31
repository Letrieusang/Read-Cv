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

package bitwise ;import (_aab "encoding/binary";_aa "errors";_a "fmt";_b "github.com/unidoc/unipdf/v3/common";_bg "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_f "io";);func (_gdc *BufferedWriter )tryGrowByReslice (_eb int )bool {if _gdf :=len (_gdc ._fe );
_eb <=cap (_gdc ._fe )-_gdf {_gdc ._fe =_gdc ._fe [:_gdf +_eb ];return true ;};return false ;};func (_dbg *Writer )UseMSB ()bool {return _dbg ._fcc };func (_gga *Reader )AbsoluteLength ()uint64 {return uint64 (len (_gga ._aed ._fae ))};type Reader struct{_aed readerSource ;
_aadd byte ;_gcf byte ;_gcea int64 ;_fac int ;_caa int ;_bbe int64 ;_ceb byte ;_age byte ;_acc int ;};type BufferedWriter struct{_fe []byte ;_bd uint8 ;_df int ;_bdb bool ;};func (_gb *BufferedWriter )Reset (){_gb ._fe =_gb ._fe [:0];_gb ._df =0;_gb ._bd =0};
func (_ec *BufferedWriter )byteCapacity ()int {_ce :=len (_ec ._fe )-_ec ._df ;if _ec ._bd !=0{_ce --;};return _ce ;};func (_gfb *Reader )readUnalignedByte ()(_dcba byte ,_daa error ){_dda :=_gfb ._gcf ;_dcba =_gfb ._aadd <<(8-_dda );_gfb ._aadd ,_daa =_gfb .readBufferByte ();
if _daa !=nil {return 0,_daa ;};_dcba |=_gfb ._aadd >>_dda ;_gfb ._aadd &=1<<_dda -1;return _dcba ,nil ;};func (_ad *BufferedWriter )fullOffset ()int {_ae :=_ad ._df ;if _ad ._bd !=0{_ae ++;};return _ae ;};func (_ded *Reader )Read (p []byte )(_ced int ,_dce error ){if _ded ._gcf ==0{return _ded .read (p );
};for ;_ced < len (p );_ced ++{if p [_ced ],_dce =_ded .readUnalignedByte ();_dce !=nil {return 0,_dce ;};};return _ced ,nil ;};func (_aged *Writer )writeByte (_cag byte )error {if _aged ._bfd > len (_aged ._fc )-1{return _f .EOF ;};if _aged ._bfd ==len (_aged ._fc )-1&&_aged ._bfgg !=0{return _f .EOF ;
};if _aged ._bfgg ==0{_aged ._fc [_aged ._bfd ]=_cag ;_aged ._bfd ++;return nil ;};if _aged ._fcc {_aged ._fc [_aged ._bfd ]|=_cag >>_aged ._bfgg ;_aged ._bfd ++;_aged ._fc [_aged ._bfd ]=byte (uint16 (_cag )<<(8-_aged ._bfgg )&0xff);}else {_aged ._fc [_aged ._bfd ]|=byte (uint16 (_cag )<<_aged ._bfgg &0xff);
_aged ._bfd ++;_aged ._fc [_aged ._bfd ]=_cag >>(8-_aged ._bfgg );};return nil ;};func (_bdg *Reader )ReadByte ()(byte ,error ){if _bdg ._gcf ==0{return _bdg .readBufferByte ();};return _bdg .readUnalignedByte ();};type BinaryWriter interface{BitWriter ;
_f .Writer ;_f .ByteWriter ;Data ()[]byte ;};type BitWriter interface{WriteBit (_dcgg int )error ;WriteBits (_ecd uint64 ,_cb int )(_ac int ,_fbf error );FinishByte ();SkipBits (_cgb int )error ;};func (_cd *BufferedWriter )SkipBits (skip int )error {if skip ==0{return nil ;
};_de :=int (_cd ._bd )+skip ;if _de >=0&&_de < 8{_cd ._bd =uint8 (_de );return nil ;};_de =int (_cd ._bd )+_cd ._df *8+skip ;if _de < 0{return _bg .Errorf ("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073","\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_dcg :=_de /8;_e :=_de %8;_cd ._bd =uint8 (_e );if _dcf :=_dcg -_cd ._df ;_dcf > 0&&len (_cd ._fe )-1< _dcg {if _cd ._bd !=0{_dcf ++;};_cd .expandIfNeeded (_dcf );};_cd ._df =_dcg ;return nil ;};type Writer struct{_fc []byte ;_bfgg uint8 ;_bfd int ;_fcc bool ;
};var _ _f .ByteWriter =&BufferedWriter {};func (_bbef *Reader )Mark (){_bbef ._bbe =_bbef ._gcea ;_bbef ._ceb =_bbef ._gcf ;_bbef ._age =_bbef ._aadd ;_bbef ._acc =_bbef ._fac ;};var _ BinaryWriter =&Writer {};func (_eee *Writer )FinishByte (){if _eee ._bfgg ==0{return ;
};_eee ._bfgg =0;_eee ._bfd ++;};func (_gbb *BufferedWriter )WriteByte (bt byte )error {if _gbb ._df > len (_gbb ._fe )-1||(_gbb ._df ==len (_gbb ._fe )-1&&_gbb ._bd !=0){_gbb .expandIfNeeded (1);};_gbb .writeByte (bt );return nil ;};func (_af *BufferedWriter )Data ()[]byte {return _af ._fe };
type StreamReader interface{_f .Reader ;_f .ByteReader ;_f .Seeker ;Align ()byte ;BitPosition ()int ;Mark ();Length ()uint64 ;ReadBit ()(int ,error );ReadBits (_ag byte )(uint64 ,error );ReadBool ()(bool ,error );ReadUint32 ()(uint32 ,error );Reset ();
AbsolutePosition ()int64 ;};func (_cdd *Reader )AbsolutePosition ()int64 {return _cdd ._gcea +int64 (_cdd ._aed ._agd )};func (_cc *Reader )ReadUint32 ()(uint32 ,error ){_adb :=make ([]byte ,4);_ ,_acb :=_cc .Read (_adb );if _acb !=nil {return 0,_acb ;
};return _aab .BigEndian .Uint32 (_adb ),nil ;};func (_bed *Writer )Data ()[]byte {return _bed ._fc };func (_db *BufferedWriter )Len ()int {return _db .byteCapacity ()};func (_aacc *BufferedWriter )WriteBit (bit int )error {if bit !=1&&bit !=0{return _bg .Errorf ("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064",bit );
};if len (_aacc ._fe )-1< _aacc ._df {_aacc .expandIfNeeded (1);};_ab :=_aacc ._bd ;if _aacc ._bdb {_ab =7-_aacc ._bd ;};_aacc ._fe [_aacc ._df ]|=byte (uint16 (bit <<_ab )&0xff);_aacc ._bd ++;if _aacc ._bd ==8{_aacc ._df ++;_aacc ._bd =0;};return nil ;
};func (_ca *BufferedWriter )Write (d []byte )(int ,error ){_ca .expandIfNeeded (len (d ));if _ca ._bd ==0{return _ca .writeFullBytes (d ),nil ;};return _ca .writeShiftedBytes (d ),nil ;};func (_dcb *BufferedWriter )WriteBits (bits uint64 ,number int )(_bad int ,_cg error ){const _gc ="\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s";
if number < 0||number > 64{return 0,_bg .Errorf (_gc ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};_be :=number /8;if _be > 0{_bgc :=number -_be *8;for _baa :=_be -1;_baa >=0;_baa --{_gf :=byte ((bits >>uint (_baa *8+_bgc ))&0xff);if _cg =_dcb .WriteByte (_gf );_cg !=nil {return _bad ,_bg .Wrapf (_cg ,_gc ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_be -_baa +1);
};};number -=_be *8;if number ==0{return _be ,nil ;};};var _da int ;for _bc :=0;_bc < number ;_bc ++{if _dcb ._bdb {_da =int ((bits >>uint (number -1-_bc ))&0x1);}else {_da =int (bits &0x1);bits >>=1;};if _cg =_dcb .WriteBit (_da );_cg !=nil {return _bad ,_bg .Wrapf (_cg ,_gc ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_bc );
};};return _be ,nil ;};type readerSource struct{_fae []byte ;_agd int ;_cbd int ;};func (_aedd *Reader )readBufferByte ()(byte ,error ){if _aedd ._gcea >=int64 (_aedd ._aed ._cbd ){return 0,_f .EOF ;};_aedd ._caa =-1;_ace :=_aedd ._aed ._fae [int64 (_aedd ._aed ._agd )+_aedd ._gcea ];
_aedd ._gcea ++;_aedd ._fac =int (_ace );return _ace ,nil ;};func (_c *BufferedWriter )ResetBitIndex (){_c ._bd =0};func (_ge *Reader )Length ()uint64 {return uint64 (_ge ._aed ._cbd )};func (_fa *BufferedWriter )writeShiftedBytes (_def []byte )int {for _ ,_bf :=range _def {_fa .writeByte (_bf );
};return len (_def );};func (_dbf *Reader )ReadBool ()(bool ,error ){return _dbf .readBool ()};func (_gbd *Reader )NewPartialReader (offset ,length int ,relative bool )(*Reader ,error ){if offset < 0{return nil ,_aa .New ("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065");
};if relative {offset =_gbd ._aed ._agd +offset ;};if length > 0{_bcg :=len (_gbd ._aed ._fae );if relative {_bcg =_gbd ._aed ._cbd ;};if offset +length > _bcg {return nil ,_a .Errorf ("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064",offset ,length ,offset +length ,_gbd ._aed ._cbd );
};};if length < 0{_bae :=len (_gbd ._aed ._fae );if relative {_bae =_gbd ._aed ._cbd ;};length =_bae -offset ;};return &Reader {_aed :readerSource {_fae :_gbd ._aed ._fae ,_cbd :length ,_agd :offset }},nil ;};func (_dcd *BufferedWriter )writeByte (_gg byte ){switch {case _dcd ._bd ==0:_dcd ._fe [_dcd ._df ]=_gg ;
_dcd ._df ++;case _dcd ._bdb :_dcd ._fe [_dcd ._df ]|=_gg >>_dcd ._bd ;_dcd ._df ++;_dcd ._fe [_dcd ._df ]=byte (uint16 (_gg )<<(8-_dcd ._bd )&0xff);default:_dcd ._fe [_dcd ._df ]|=byte (uint16 (_gg )<<_dcd ._bd &0xff);_dcd ._df ++;_dcd ._fe [_dcd ._df ]=_gg >>(8-_dcd ._bd );
};};func NewReader (data []byte )*Reader {return &Reader {_aed :readerSource {_fae :data ,_cbd :len (data ),_agd :0}};};func (_bga *Reader )Align ()(_abb byte ){_abb =_bga ._gcf ;_bga ._gcf =0;return _abb };func (_abd *Writer )byteCapacity ()int {_fed :=len (_abd ._fc )-_abd ._bfd ;
if _abd ._bfgg !=0{_fed --;};return _fed ;};func (_bce *Writer )ResetBit (){_bce ._bfgg =0};func (_ea *BufferedWriter )expandIfNeeded (_gd int ){if !_ea .tryGrowByReslice (_gd ){_ea .grow (_gd );};};func NewWriter (data []byte )*Writer {return &Writer {_fc :data }};
var (_ _f .Reader =&Reader {};_ _f .ByteReader =&Reader {};_ _f .Seeker =&Reader {};_ StreamReader =&Reader {};);func BufferedMSB ()*BufferedWriter {return &BufferedWriter {_bdb :true }};func (_feb *BufferedWriter )grow (_gce int ){if _feb ._fe ==nil &&_gce < _aac {_feb ._fe =make ([]byte ,_gce ,_aac );
return ;};_fg :=len (_feb ._fe );if _feb ._bd !=0{_fg ++;};_bb :=cap (_feb ._fe );switch {case _gce <=_bb /2-_fg :_b .Log .Trace ("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027",len (_feb ._fe ),cap (_feb ._fe ),_gce );
_b .Log .Trace ("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027",_bb ,_fg );copy (_feb ._fe ,_feb ._fe [_feb .fullOffset ():]);
case _bb > _ba -_bb -_gce :_b .Log .Error ("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065");return ;default:_cdf :=make ([]byte ,2*_bb +_gce );copy (_cdf ,_feb ._fe );_feb ._fe =_cdf ;};_feb ._fe =_feb ._fe [:_fg +_gce ];
};func (_add *Reader )ReadBits (n byte )(_agdg uint64 ,_bef error ){if n < _add ._gcf {_bfg :=_add ._gcf -n ;_agdg =uint64 (_add ._aadd >>_bfg );_add ._aadd &=1<<_bfg -1;_add ._gcf =_bfg ;return _agdg ,nil ;};if n > _add ._gcf {if _add ._gcf > 0{_agdg =uint64 (_add ._aadd );
n -=_add ._gcf ;};for n >=8{_dbc ,_bfe :=_add .readBufferByte ();if _bfe !=nil {return 0,_bfe ;};_agdg =_agdg <<8+uint64 (_dbc );n -=8;};if n > 0{if _add ._aadd ,_bef =_add .readBufferByte ();_bef !=nil {return 0,_bef ;};_bbc :=8-n ;_agdg =_agdg <<n +uint64 (_add ._aadd >>_bbc );
_add ._aadd &=1<<_bbc -1;_add ._gcf =_bbc ;}else {_add ._gcf =0;};return _agdg ,nil ;};_add ._gcf =0;return uint64 (_add ._aadd ),nil ;};func (_cf *Reader )readBool ()(_ee bool ,_eea error ){if _cf ._gcf ==0{_cf ._aadd ,_eea =_cf .readBufferByte ();if _eea !=nil {return false ,_eea ;
};_ee =(_cf ._aadd &0x80)!=0;_cf ._aadd ,_cf ._gcf =_cf ._aadd &0x7f,7;return _ee ,nil ;};_cf ._gcf --;_ee =(_cf ._aadd &(1<<_cf ._gcf ))!=0;_cf ._aadd &=1<<_cf ._gcf -1;return _ee ,nil ;};func (_gca *Reader )read (_ga []byte )(int ,error ){if _gca ._gcea >=int64 (_gca ._aed ._cbd ){return 0,_f .EOF ;
};_gca ._caa =-1;_bgf :=copy (_ga ,_gca ._aed ._fae [(int64 (_gca ._aed ._agd )+_gca ._gcea ):(_gca ._aed ._agd +_gca ._aed ._cbd )]);_gca ._gcea +=int64 (_bgf );return _bgf ,nil ;};func (_eba *Reader )ReadBit ()(_acf int ,_cgg error ){_cda ,_cgg :=_eba .readBool ();
if _cgg !=nil {return 0,_cgg ;};if _cda {_acf =1;};return _acf ,nil ;};func (_fba *Reader )RelativePosition ()int64 {return _fba ._gcea };func (_dbcb *Writer )writeBit (_acca uint8 )error {if len (_dbcb ._fc )-1< _dbcb ._bfd {return _f .EOF ;};_ef :=_dbcb ._bfgg ;
if _dbcb ._fcc {_ef =7-_dbcb ._bfgg ;};_dbcb ._fc [_dbcb ._bfd ]|=byte (uint16 (_acca <<_ef )&0xff);_dbcb ._bfgg ++;if _dbcb ._bfgg ==8{_dbcb ._bfd ++;_dbcb ._bfgg =0;};return nil ;};func (_dcec *Writer )WriteBit (bit int )error {switch bit {case 0,1:return _dcec .writeBit (uint8 (bit ));
};return _bg .Error ("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065");};func (_bca *Writer )Write (p []byte )(int ,error ){if len (p )> _bca .byteCapacity (){return 0,_f .EOF ;
};for _ ,_aff :=range p {if _dfc :=_bca .writeByte (_aff );_dfc !=nil {return 0,_dfc ;};};return len (p ),nil ;};func (_ddf *Reader )Seek (offset int64 ,whence int )(int64 ,error ){_ddf ._caa =-1;_ddf ._gcf =0;_ddf ._aadd =0;_ddf ._fac =0;var _ece int64 ;
switch whence {case _f .SeekStart :_ece =offset ;case _f .SeekCurrent :_ece =_ddf ._gcea +offset ;case _f .SeekEnd :_ece =int64 (_ddf ._aed ._cbd )+offset ;default:return 0,_aa .New ("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065");
};if _ece < 0{return 0,_aa .New ("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e");};_ddf ._gcea =_ece ;
_ddf ._gcf =0;return _ece ,nil ;};func (_bcgc *Reader )BitPosition ()int {return int (_bcgc ._gcf )};func (_g *BufferedWriter )FinishByte (){if _g ._bd ==0{return ;};_g ._bd =0;_g ._df ++;};var _ BinaryWriter =&BufferedWriter {};func (_cfd *Writer )WriteByte (c byte )error {return _cfd .writeByte (c )};
func (_fb *BufferedWriter )writeFullBytes (_aad []byte )int {_bee :=copy (_fb ._fe [_fb .fullOffset ():],_aad );_fb ._df +=_bee ;return _bee ;};func (_gde *Reader )ConsumeRemainingBits ()(uint64 ,error ){if _gde ._gcf !=0{return _gde .ReadBits (_gde ._gcf );
};return 0,nil ;};func (_dd *Reader )Reset (){_dd ._gcea =_dd ._bbe ;_dd ._gcf =_dd ._ceb ;_dd ._aadd =_dd ._age ;_dd ._fac =_dd ._acc ;};func (_gab *Writer )WriteBits (bits uint64 ,number int )(_bcaa int ,_cfa error ){const _afg ="\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073";
if number < 0||number > 64{return 0,_bg .Errorf (_afg ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};if number ==0{return 0,nil ;};_cac :=number /8;if _cac > 0{_ead :=number -_cac *8;for _geg :=_cac -1;_geg >=0;_geg --{_ged :=byte ((bits >>uint (_geg *8+_ead ))&0xff);if _cfa =_gab .WriteByte (_ged );_cfa !=nil {return _bcaa ,_bg .Wrapf (_cfa ,_afg ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_cac -_geg +1);
};};number -=_cac *8;if number ==0{return _cac ,nil ;};};var _eae int ;for _ecg :=0;_ecg < number ;_ecg ++{if _gab ._fcc {_eae =int ((bits >>uint (number -1-_ecg ))&0x1);}else {_eae =int (bits &0x1);bits >>=1;};if _cfa =_gab .WriteBit (_eae );_cfa !=nil {return _bcaa ,_bg .Wrapf (_cfa ,_afg ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_ecg );
};};return _cac ,nil ;};func NewWriterMSB (data []byte )*Writer {return &Writer {_fc :data ,_fcc :true }};const (_aac =64;_ba =int (^uint (0)>>1););var _ _f .Writer =&BufferedWriter {};func (_gff *Writer )SkipBits (skip int )error {const _ed ="\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073";
if skip ==0{return nil ;};_ebc :=int (_gff ._bfgg )+skip ;if _ebc >=0&&_ebc < 8{_gff ._bfgg =uint8 (_ebc );return nil ;};_ebc =int (_gff ._bfgg )+_gff ._bfd *8+skip ;if _ebc < 0{return _bg .Errorf (_ed ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_eaf :=_ebc /8;_bba :=_ebc %8;_b .Log .Trace ("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073");_b .Log .Trace ("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027",_gff ._bfgg ,_gff ._bfd ,int (_gff ._bfgg )+(_gff ._bfd )*8,len (_gff ._fc ),cap (_gff ._fc ));
_b .Log .Trace ("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027",skip ,_ebc ,_bba );_gff ._bfgg =uint8 (_bba );if _bfgb :=_eaf -_gff ._bfd ;
_bfgb > 0&&len (_gff ._fc )-1< _eaf {_b .Log .Trace ("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064",_bfgb );return _bg .Errorf (_ed ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_gff ._bfd =_eaf ;
_b .Log .Trace ("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027",_gff ._bfgg ,_gff ._bfd );return nil ;};