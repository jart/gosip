# RFC4475 Torture Message String Literal Encoder

import re
import sys
import textwrap

hexsub = lambda m: re.sub(r'(..)', r'\x\1', m.group(1))
repeatsub = lambda m: m.group(2) * int(m.group(1))
allonelinesub = lambda m: m.group(1).replace('\n', '')

message = sys.stdin.read()
message = textwrap.dedent(message)
message = re.sub(r'<allOneLine>(.+?)</allOneLine>', allonelinesub, message, 0, re.S)
message = re.sub(r'<repeat count=(\d+)>(.*?)</repeat>', repeatsub, message, 0, re.S)
message = message.replace('\\', '\\\\')
message = re.sub(r'<hex>(.*?)</hex>', hexsub, message, 0, re.S)
message = message.replace('"', '\\"')
message = '"' + message.replace('\n', '\\r\\n" +\n"') + '"'
print message
