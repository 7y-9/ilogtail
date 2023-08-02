#include "MetricExportor.h"
#include "sender/Sender.h"
#include "log_pb/sls_logs.pb.h"
#include "common/LogtailCommonFlags.h"

using namespace sls_logs;
using namespace std;

namespace logtail {

MetricExportor::MetricExportor() {
    mSendInterval = 10;
    mSnapshotInterval = 10;
    mLastSendTime = 0;
    mLastSnapshotTime = 0;
}

void MetricExportor::PushMetrics() {
    ReadMetrics::GetInstance()->UpdateMetrics();
    sls_logs::LogGroup logGroup;
    ReadMetrics::GetInstance()->ReadAsLogGroup(logGroup);
    // TODO: send
}
}