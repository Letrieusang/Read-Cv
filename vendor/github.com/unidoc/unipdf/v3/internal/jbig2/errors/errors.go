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

package errors ;import (_g "fmt";_f "golang.org/x/xerrors";);func Error (processName ,message string )error {return _db (message ,processName )};func (_d *processError )Error ()string {var _bc string ;if _d ._ge !=""{_bc =_d ._ge ;};_bc +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_d ._fg ;
if _d ._b !=""{_bc +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_d ._b ;};if _d ._ag !=nil {_bc +="\u002e\u0020"+_d ._ag .Error ();};return _bc ;};type processError struct{_ge string ;_fg string ;_b string ;_ag error ;};func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _ed ,_ged :=err .(*processError );
_ged {_ed ._ge ="";};_af :=_db (_g .Sprintf (message ,arguments ...),processName );_af ._ag =err ;return _af ;};func (_ab *processError )Unwrap ()error {return _ab ._ag };func Wrap (err error ,processName ,message string )error {if _bd ,_ga :=err .(*processError );
_ga {_bd ._ge ="";};_dd :=_db (message ,processName );_dd ._ag =err ;return _dd ;};var _ _f .Wrapper =(*processError )(nil );func _db (_c ,_bf string )*processError {return &processError {_ge :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_b :_c ,_fg :_bf };
};func Errorf (processName ,message string ,arguments ...interface{})error {return _db (_g .Sprintf (message ,arguments ...),processName );};