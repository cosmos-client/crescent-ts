import * as ununificlient from './module';
import Long from 'long';
import * as $protobuf from 'protobufjs/minimal';

$protobuf.util.Long = Long;
$protobuf.configure();

export default ununificlient;
