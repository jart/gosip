# Copyright 2020 Justine Alexandra Roberts Tunney
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
